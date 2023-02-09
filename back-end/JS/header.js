function loadHeader() {
  // fetch the header code
  fetch("header.html")
    .then((response) => response.text())
    .then((header) => {
      // insert the header code into the header placeholder
      document.querySelector("#header-placeholder").innerHTML = header;
      // load the signup script
      loadSignUp();
    });
}

function loadSignUp() {
  const signupBtn = document.querySelector("#signup-btn");
  const signupFormModal = document.querySelector("#signup-form-modal");
  const closeBtn = document.querySelector(".close-btn");

  signupBtn.addEventListener("click", function () {
    signupFormModal.style.display = "block";
  });

  closeBtn.addEventListener("click", function () {
    signupFormModal.style.display = "none";
  });
}
