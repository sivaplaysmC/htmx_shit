console.warn("Hi from auth_helper!!!")


// Wrap the fetch function to intercept and modify requests
function fetchWithAuthorization(url, options = {}) {
  const modifiedOptions = addAuthorizationToken(url, options);
  options.headers = {
    ...options.headers,
    Authorization: `Bearer siva`
  }
  return fetch(url, modifiedOptions);
}
window.fetch = fetchWithAuthorization;


document.addEventListener("htmx:configRequest", event => {
  event.detail.headers["Authorization"] = "Bearer siva"
})
