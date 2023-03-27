class Post extends HTMLElement {
  constructor(
    ID,
    category,
    title,
    content,
    date,
    comments,
    rating,
    userID,
    username,
    userAvatar
  ) {
    super();

    this.ID = ID;
    this.category = category;
    this.title = title;
    this.content = content;
    this.date = date;
    this.comments = comments;
    this.rating = rating;
    this.userID = userID;
    this.username = username;
    this.userAvatar = userAvatar;

    // Create HTML elements and set values
    this.dataset.id = this.ID;
    this.classList.add("post");
    this.innerHTML = `
    <header class="post__header">
    <a class="post__username" href="#/profile/${this.userID}">${this.username}</a>
    <span class="post__separator">|</span>
    <div class="post__title">${this.title}</div>
    <time class="post__date">${this.date}</time>
    <div class="post__comments">
      <img
        class="post__comments--icon"
        src="images/comment.svg"
        draggable="false"
        alt="Comments"
      />
      <b class="post__comments--count">${this.comments.length}</b>
    </div>
    <div class="post__rating">
      <img
        class="post__rating--icon"
        src="images/face-smile.svg"
        draggable="false"
        alt="Comments"
      />
      <b class="post__rating--count">${this.rating}</b>
    </div>
  </header>
  <section class="post__body">
    <div class="post__profile">
      <img
        class="post__avatar"
        src="${this.userAvatar}"
        alt="${this.username}'s avatar"
      />
    </div>
    <div class="post__content">
      ${this.content}
    </div>
    <div class="post__interactions">
      <img
        class="post__interaction post__interaction--upvote"
        src="images/arrow-up.svg"
        alt="Upvote"
        draggable="false"
      />
      <img
        class="post__interaction post__interaction--downvote"
        src="images/arrow-down.svg"
        alt="Downvote"
        draggable="false"
      />
      <img
        class="post__interaction post__interaction--comment"
        src="images/comment.svg"
        alt="Comment"
        draggable="false"
      />
    </div>
  </section>
  <section class="post__comments"></section>`;

    // Append post to container
    const container = document.querySelector(".post-container");
    container.appendChild(this);
  }
}
customElements.define("post-element", Post);

// Fetch all posts from database endpoint and append to page
function populatePosts() {
  fetch("/api/posts")
    .then((response) => response.json())
    .then((data) => {
      data.forEach((post) => {
        let postElement = new Post(
          post.ID,
          post.category,
          post.title,
          post.content,
          post.date,
          post.comments,
          post.rating,
          post.userID,
          post.username,
          post.userAvatar
        );
        // Make post expandable
        postElement
          .querySelector(".post__header")
          .addEventListener("click", (event) => {
            if (!event.target.classList.contains("post__username")) {
              postElement.classList.toggle("post--expanded");
            }
          });
        // Add event listeners to post interactions
        postElement
          .querySelector(".post__interaction--upvote")
          .addEventListener("click", () => this.rating++);
        postElement
          .querySelector(".post__interaction--downvote")
          .addEventListener("click", () => this.rating--);
        postElement
          .querySelector(".post__interaction--comment")
          .addEventListener("click", () => this.comments++);

        // Add comments to post
        post.comments.forEach((comment) => {
          let commentElement = new Post(
            comment.ID,
            comment.content,
            comment.date,
            comment.rating,
            comment.userID,
            comment.username,
            comment.userAvatar
          );
          postElement
            .querySelector(".post__comments")
            .appendChild(commentElement);
        });
      });
    })
    .catch((error) => console.error(error));
}

// DOM loaded
document.addEventListener("DOMContentLoaded", () => {
  populatePosts();
});
