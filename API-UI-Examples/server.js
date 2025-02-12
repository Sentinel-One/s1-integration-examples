// server.js
const express = require('express');
const bodyParser = require('body-parser');
const axios = require('axios');
const session = require('express-session');

const app = express();
const port = 3000;

app.use(bodyParser.json());
app.use(express.static('public')); 
app.use(session({ 
  secret: 's1_console_light_sess_key', 
  resave: false, 
  saveUninitialized: true 
}));

//-----------------------------------------
// URLs
//-----------------------------------------
app.get('/', (req, res) => {
    res.sendFile(__dirname + '/public/main.html'); 
});

app.get('/logout', (req, res) => {
    // Remove Session settings
    req.session.destroy(err => {
        if (err) {
          console.error(err);
          res.status(500).send('Error destroying session');
        } 
      });
    res.redirect('/');
});

app.post('/login', async (req, res) => {
    var { apiURL, apiToken } = req.body;

    // Remove Session settings
    req.session.isAuthenticated = false;
    req.session.apiURL = "";
    req.session.apiToken = "";

    // Validate URL & Token with an API Call
    try {
      const response =  await axios({
        method: 'GET',
        url: `${apiURL}/web/api/v2.1/accounts`,
        headers: {
          'Authorization': `ApiToken ${apiToken}`
        }
      });

      if (response.status === 200) { // Assuming successful validation returns 200 OK
        req.session.isAuthenticated = true;
        req.session.apiURL = apiURL;
        req.session.apiToken = apiToken;
        res.redirect('/');
      } else {
        //res.redirect('/'); // Or display an error message
        res.status(response.status).json({ error: 'Invalid URL or Token.  Error:' + response.status });
      }
    } catch (error) {
      console.error('Authentication failed:', error);
      res.status(400).json({ error: 'Invalid URL or Token.' });
      //res.redirect('/xxx'); // Or display an error message
    }
});

app.post('/execute', async (req, res) => {
  const { endpoint } = req.body;
  const apiURL = req.session.apiURL + '/web/api/v2.1';
  const apiToken = req.session.apiToken;

  try {
    const response = await axios({
      method: 'GET',
      url: `${apiURL}${endpoint}`,
      headers: {
        'Authorization': `ApiToken ${apiToken}`
      }
    });

    res.json(response.data);
  } catch (error) {
    console.error('Error executing API call:', error);
    res.status(500).json({ error: 'Failed to execute API call' });
  }
});

app.listen(port, () => {
  console.log(`Server listening on port ${port}`);
});
