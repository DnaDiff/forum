document.addEventListener("DOMContentLoaded", function () {
  let loginForm = document.getElementById("login-form");

  loginForm.addEventListener("submit", function (event) {
    event.preventDefault();

    // get form data
    let username = loginForm.querySelector('input[type="text"]').value;
    let password = loginForm.querySelector('input[type="password"]').value;

    // create user object
    let user = {
      username: username,
      password: password,
    };

    // send POST request to /login endpoint
    fetch("http://localhost:8080/api/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(user),
    })
      .then((response) => response.text())
      .then((data) => console.log(data))
      .catch((error) => console.error("Error:", error));
  });
});
