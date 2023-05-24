let registerForm = document.getElementById("register-form");

let elements = document.getElementsByClassName("form-input");

async function submitForm(user) {
  console.log(user);
  // send POST request to /register endpoint
  try {
    let response = await fetch("http://localhost:8080/api/register", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(user),
    });
    let data = await response.text();
    console.log(data);
  } catch (error) {
    console.error("Error:", error);
  }
}

function getUserData() {
  // get form data
  return {
    username: registerForm.querySelector('input[name="username"]').value,
    email: registerForm.querySelector('input[name="email"]').value,
    firstname: registerForm.querySelector('input[name="firstname"]').value,
    lastname: registerForm.querySelector('input[name="lastname"]').value,
    gender: registerForm.querySelector('select[name="gender"]').value,
    age: registerForm.querySelector('input[name="age"]').value,
    password: registerForm.querySelector('input[name="password"]').value,
    confirmPassword: registerForm.querySelector('input[name="confirmpassword"]')
      .value,
  };
}

registerForm.addEventListener("submit", function (event) {
  event.preventDefault();

  let user = getUserData();
  submitForm(user);
});
