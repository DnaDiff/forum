class Post extends HTMLElement {
  constructor(
    ID,
    parentID,
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
    this.parentID = parentID;
    this.title = title;
    this.content = content;
    this.date = date;
    this.comments = comments.length;
    this.rating = rating;
    this.userID = userID;
    this.username = username;
    this.userAvatar = userAvatar;

    // Append post to container
    document.querySelector(".post-container").appendChild(this); // Replace with category once implemented
  }

  connectedCallback() {
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
      <b class="post__comments--count">${this.comments}</b>
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

    // Make post expandable
    this.querySelector(".post__header").addEventListener("click", (event) => {
      if (!event.target.classList.contains("post__username")) {
        this.classList.toggle("post--expanded");
      }
    });
    // Add event listeners to post interactions
    addInteractionListener(
      this.querySelector(".post__interaction--upvote"),
      `/api/posts/${ID}/upvote`,
      "POST",
      (data) => {
        this.querySelector(".post__rating--count").textContent = data.rating;
      }
    );
    addInteractionListener(
      this.querySelector(".post__interaction--downvote"),
      `/api/posts/${ID}/downvote`,
      "PUT",
      (data) => {
        this.querySelector(".post__rating--count").textContent = data.rating;
      }
    );
    addInteractionListener(
      this.querySelector(".post__interaction--comment"),
      `/api/posts/${ID}/comment`,
      "POST",
      (data) => {
        console.log("Comment");
        // this.querySelector(".post__comments--count").textContent = data.comments.length;
      }
    );
  }

  disconnectedCallback() {
    console.log("Post removed");

    // Remove event listeners
    this.querySelector(".post__header").removeEventListener();
    this.querySelector(".post__interaction--upvote").removeEventListener();
    this.querySelector(".post__interaction--downvote").removeEventListener();
    this.querySelector(".post__interaction--comment").removeEventListener();

    // Remove post
    this.remove();
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
          post.parentID,
          post.title,
          post.content,
          post.date,
          post.comments,
          post.rating,
          post.userID,
          post.username,
          post.userAvatar
        );

        // Add comments to post
        post.comments.forEach((comment) => {
          let commentElement = new Post(
            comment.ID,
            comment.parentID,
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

function addInteractionListener(element, endpoint, method, callback) {
  element.addEventListener("click", () => {
    fetch(endpoint, {
      method: method,
    })
      .then((response) => response.json())
      .then((data) => {
        callback(data);
      })
      .catch((error) => console.error("Interaction failed:", error));
  });
}

// DOM loaded
document.addEventListener("DOMContentLoaded", () => {
  populatePosts();
});
