
# HeaderHasher Traefik middleware

This is a simple middleware that looks for a specific HTTP request header and if found attaches its sha256 value as another header. 

An example practical application of this would custom rate limiting routes for specific authorisation headers. You don't want to keep the plaintext
authorization token value together with the rest of ingress routes configuration. This middleware allows to safely use it's hash instead keeping the token secret.

## Usage

TODO

### Configuration

TODO
