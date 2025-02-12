package scanner

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	// scanner verdicts
	VerdictBenign     = iota // the file is benign
	VerdictSuspicious        // the file is suspicious
	VerdictMalicious         // the file is malicious
	VerdictUnknown           // the scanner was unable to determine a verdict or failed before scanning
	VerdictError             // an error occurred while actually scanning the file
	VerdictSkipped           // the file was not scanned
	VerdictRejected          // the file was rejected
)

var (
	_VerdictToStr = map[Verdict]string{
		VerdictBenign:     "benign",
		VerdictSuspicious: "suspicious",
		VerdictMalicious:  "malicious",
		VerdictUnknown:    "unknown",
		VerdictError:      "error",
		VerdictSkipped:    "skipped",
		VerdictRejected:   "rejected",
	}
	_strToVerdict = map[string]Verdict{}
)

func init() {
	for verdict, str := range _VerdictToStr {
		_strToVerdict[str] = verdict
	}
}

// Verdict holds the actual verdict of a file scan.
type Verdict int

// ParseVerdict converts the given string into a [Verdict].
func ParseVerdict(v string) (Verdict, error) {
	if verdict, ok := _strToVerdict[strings.ToLower(v)]; ok {
		return verdict, nil
	}
	return VerdictError, fmt.Errorf(" %s: unsupported scan verdict", v)
}

// MarshalJSON returns the verdict as a string formatted as JSON.
func (v Verdict) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.String())
}

// MarshalText returns the verdict as a string formatted in plain text.
func (v Verdict) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

// String returns the scan verdict as a string.
func (v Verdict) String() string {
	if str, ok := _VerdictToStr[v]; ok {
		return str
	}

	// this should never happen unless someone explicitly sets the value of a verdict to something unsupported
	return fmt.Sprintf("%d: unsupported scan verdict value", v)
}

// UnmarshalJSON parses the given JSON-formatted data into a [Verdict] object.
func (v *Verdict) UnmarshalJSON(data []byte) error {
	var verdict int
	if err := json.Unmarshal(data, &verdict); err == nil {
		if _, ok := _VerdictToStr[Verdict(verdict)]; ok {
			*v = Verdict(verdict)
			return nil
		}
		return fmt.Errorf("%d: unsupported scan verdict value", verdict)
	}

	var strVerdict string
	if err := json.Unmarshal(data, &strVerdict); err != nil {
		return err
	}
	parsedVerdict, err := ParseVerdict(strVerdict)
	if err != nil {
		return err
	}
	*v = parsedVerdict
	return nil
}

// UnmarshalText parses the given plain text data into a [Verdict] object.
func (v *Verdict) UnmarshalText(data []byte) error {
	parsedVerdict, err := ParseVerdict(string(data))
	if err != nil {
		return err
	}
	*v = parsedVerdict
	return nil
}
