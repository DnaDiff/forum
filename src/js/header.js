function loadHeader() {
  // fetch the header code
  fetch("../../templates/header.html")
    .then((response) => response.text())
    .then((header) => {
      // insert the header code into the header placeholder
      document.querySelector("#header-placeholder").innerHTML = header;
      // load the signup script
      loadSignUp();
      loadSignIn();
      loadPost();
      loadComment();
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

function loadPost() {
  const dimmer = document.querySelector("#dim");
  const postBtn = document.querySelector("#post-btn");
  const postFormModal = document.querySelector("#post-form-modal");
  const closeBtn = document.querySelector(".close-btn2");

  document.getElementById("post").addEventListener("input", updateCharCount);

  postBtn.addEventListener("click", function () {
    postFormModal.style.display = "block";
    dimmer.style.display = "block";
  });

  closeBtn.addEventListener("click", function () {
    postFormModal.style.display = "none";
    dimmer.style.display = "none";
  });
}

function loadComment() {
  const dimmer = document.querySelector("#dim");
  const commentBtn = document.querySelector("#comment-btn");
  const commentFormModal = document.querySelector("#comment-form-modal");
  const closeBtn = document.querySelector(".close-btn3");

  commentBtn.addEventListener("click", function () {
    commentFormModal.style.display = "block";
    dimmer.style.display = "block";
  });

  closeBtn.addEventListener("click", function () {
    commentFormModal.style.display = "none";
    dimmer.style.display = "none";
  });
}

function updateCharCount() {
  var content = document.getElementById("post");
  var charsLeft = document.getElementById("charsLeft");
  var maxChars = content.getAttribute("maxlength");
  var remainingChars = maxChars - content.value.length;
  charsLeft.innerHTML = remainingChars;
}
