function appendNewPost(category, title, content, author, date, avatarUrl) {
  // create a new post element
  const postElement = document.createElement("div");
  postElement.classList.add("post");
  postElement.innerHTML = `
    <div class="post-avatar">
      <img src="${"https://st3.depositphotos.com/6672868/13701/v/600/depositphotos_137014128-stock-illustration-user-profile-icon.jpg"}" alt="${author}'s avatar">
    </div>
    <div class="post-content">
      <h2>${title}</h2>
      <p class="content">${content}</p>
      <p class="author">Started by ${author} on ${date}</p>
    </div>
  `;

  // select the container element
  const categoryElement = document.getElementById(category);

  // append the new post element to the container element
  categoryElement.appendChild(postElement);
}

// Make posts expandable
document.querySelectorAll(".post__header").forEach((header) => {
  header.addEventListener("click", () => {
    const body = header.nextElementSibling;
    body.classList.toggle("post__body--expanded");
  });
});

// for (let i = 0; i < 10; i++) {
//   appendNewPost(
//     "post-container",
//     "New Post",
//     "This is a new post",
//     "John Doe" + i,
//     "2020-01-01"
//   );
// }
