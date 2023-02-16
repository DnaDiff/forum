function loadHeader() {
  // fetch the header code
  fetch("header.html")
    .then((response) => response.text())
    .then((header) => {
      // insert the header code into the header placeholder
      document.querySelector("#header-placeholder").innerHTML = header;
      // load the signup script
      loadSignUp();
      loadSignIn();
    });
}

function loadSignUp() {

  
  const dimmer = document.querySelector("#dim");
  const signupBtn = document.querySelector("#signup-btn");
  const signupFormModal = document.querySelector("#signup-form-modal");
  const closeBtn = document.querySelector(".close-btn");

  signupBtn.addEventListener("click", function () {
    signupFormModal.style.display = "block";
    dimmer.style.display = "block";
  });

  closeBtn.addEventListener("click", function () {
    signupFormModal.style.display = "none";
    dimmer.style.display = "none";
  });
}
// Reusing the same code for the sign up modal because of problems with the sign in modal not working properly when the code is not separated
function loadSignIn() {
  const dimmer = document.querySelector("#dim");
  const signinBtn = document.querySelector("#signin-btn");
  const signinFormModal = document.querySelector("#signin-form-modal");
  const closeBtn = document.querySelector(".close-btn1");

  signinBtn.addEventListener("click", function () {
    signinFormModal.style.display = "block";
    dimmer.style.display = "block";
  });

  closeBtn.addEventListener("click", function () {
    signinFormModal.style.display = "none";
    dimmer.style.display = "none";
  });
}
