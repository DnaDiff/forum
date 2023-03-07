function HeaderFunctions() {
  loadRegister();
  loadLogin();
  // loadPost();
  // loadComment();
}

function loadRegister() {
const dimmer = document.querySelector("#dim");
const registerBtn = document.querySelector("#register-btn");
const registerFormModal = document.querySelector("#register-form-modal");
const closeBtn = document.querySelector(".close-btn");

registerBtn.addEventListener("click", function () {
  registerFormModal.style.display = "block";
  dimmer.style.display = "block";
});

closeBtn.addEventListener("click", function () {
  registerFormModal.style.display = "none";
  dimmer.style.display = "none";
});
}
// Reusing the same code for the sign up modal because of problems with the sign in modal not working properly when the code is not separated
function loadLogin() {
const dimmer = document.querySelector("#dim");
const loginBtn = document.querySelector("#login-btn");
const loginFormModal = document.querySelector("#login-form-modal");
const closeBtn = document.querySelector(".close-btn1");

loginBtn.addEventListener("click", function () {
  loginFormModal.style.display = "block";
  dimmer.style.display = "block";
});

closeBtn.addEventListener("click", function () {
  loginFormModal.style.display = "none";
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

HeaderFunctions();
console.log("Header.js loaded");