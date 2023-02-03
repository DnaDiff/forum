function loadHeader() {
  // fetch the header code
  fetch("header.html")
    .then((response) => response.text())
    .then((header) => {
      // insert the header code into the header placeholder
      document.querySelector("#header-placeholder").innerHTML = header;
    });
}

const hamburger = document.querySelector(".hamburger");
const dropdown = document.querySelector(".dropdown");

hamburger.addEventListener("click", () => {
  dropdown.style.display =
    dropdown.style.display === "block" ? "none" : "block";
});
