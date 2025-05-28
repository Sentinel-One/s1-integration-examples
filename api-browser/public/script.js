//-----------------------------------------
// Components
//-----------------------------------------
const loginForm = document.getElementById('login-form');
const loginSection = document.getElementById('login-section');
const apiSection = document.getElementById('api-section');
const logoutForm = document.getElementById('logout-form');
const outputSection = document.getElementById('response-output');
const apiButtons = document.querySelectorAll('.api-button');

//-----------------------------------------
// Event listeners  - Login, Logout & API calls
//-----------------------------------------
logoutForm.addEventListener('submit', async (event) => {
    event.preventDefault();
    loginSection.style.display = 'block';
    apiSection.style.display = 'none';
    await fetch('/logout');
});

loginForm.addEventListener('submit', async (event) => {
  event.preventDefault();

  const apiURL = document.getElementById('apiURL').value;
  const apiToken = document.getElementById('apiToken').value;

  try {
    const response =  await fetch('/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ apiURL, apiToken })
    });

    if (response.ok) {
      loginSection.style.display = 'none';
      apiSection.style.display = 'block';
    } else {
      alert('Login failed. Please check your credentials.   Error: '+ response.status);
    }
  } catch (error) {
    console.error('Error during login:', error);
    alert('Invalid URL or Token.');
  }
});


apiButtons.forEach(button => {
  button.addEventListener('click', () => {
    const endpoint = button.getAttribute('data-endpoint');

    fetch('/execute', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ endpoint })
      })
    .then(response => response.json())
    .then(data => {
        // Render the JSON data
        outputSection.innerHTML = ''; // Clear existing content
        //outputSection.textContent = JSON.stringify(data, null, 2);
        $(outputSection).jsonViewer(data, {collapsed: false});
        document.getElementById('apiName').innerHTML = '<h2>/web/api/v2.1' + endpoint + '</h2>';
    })
    .catch(error => {
        console.error('Error fetching JSON data:', error);
        document.getElementById('response-output').textContent = 'Error executing API call.';
    });
  });
});

// Initially show the login section
loginSection.style.display = 'block';
apiSection.style.display = 'none';