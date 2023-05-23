let registerForm = document.getElementById("register-form");

registerForm.addEventListener("submit", function (event) {
  event.preventDefault();

  // get form data
  let username = registerForm.querySelector('input[type="text"]').value;
  let email = registerForm.querySelector('input[type="email"]').value;
  let password = registerForm.querySelector(
    'input[type="password"]:nth-child(1)'
  ).value;
  //   let confirmPassword = registerForm.querySelector(
  //     'input[type="password"]:nth-child(2)'
  //   ).value;

  // create user object
  let user = {
    username: username,
    email: email,
    password: password,
    // confirmPassword: confirmPassword,
  };

  console.log(user);

  // send POST request to /register endpoint
  fetch("http://localhost:8080/api/register", {
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
