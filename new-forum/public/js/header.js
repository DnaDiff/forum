document.addEventListener("DOMContentLoaded", function () {
  loadButtons();
  console.log("DOM loaded");
});

function loadButtons() {
  const MODAL_BUTTONS = document.querySelectorAll(".modal-btn");
  const DIM = document.querySelector("#dim");

  MODAL_BUTTONS.forEach((button) => {
    const MODAL = document.querySelector(`#${button.dataset.modal}`);
    const CLOSE_BUTTON = MODAL.querySelector(".close-btn");

    function toggleModal(event) {
      if (event.target === DIM || event.target === CLOSE_BUTTON) {
        MODAL.style.display = "none";
        DIM.style.display = "none";
      } else {
        MODAL.style.display = "block";
        DIM.style.display = "block";
      }
    }

    button.addEventListener("click", toggleModal);
    CLOSE_BUTTON.addEventListener("click", toggleModal);
    DIM.addEventListener("click", toggleModal);
  });
}

function loadPost() {
  const dimmer = document.querySelector("#dim");
  const postBtn = document.querySelector("#post-btn");
  const postFormModal = document.querySelector("#post-form-modal");
  const closeBtn = document.querySelector(".close-btn");

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

function updateCharCount() {
  var content = document.getElementById("post");
  var charsLeft = document.getElementById("charsLeft");
  var maxChars = content.getAttribute("maxlength");
  var remainingChars = maxChars - content.value.length;
  charsLeft.innerHTML = remainingChars;
}
