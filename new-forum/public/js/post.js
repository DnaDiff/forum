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

    this._ID = ID;
    this._parentID = parentID;
    this._title = title;
    this._content = content;
    this._date = date;
    this._comments = comments.length;
    this._rating = rating;
    this._userID = userID;
    this._username = username;
    this._userAvatar = userAvatar;

    // Append post to category
    if (this._parentID === null) {
      throw new Error("Post must have a parentID");
    }
    document
      .querySelector(`.post-category[id="${this._parentID}"]`)
      .appendChild(this);
  }

  connectedCallback() {
    // Create HTML elements and set values
    this.id = this._ID;
    this.classList.add("post");
    this.innerHTML = `
    <header class="post__header">
    <a class="post__username" href="#/profile/${this._userID}">${this._username}</a>
    <span class="post__separator">|</span>
    <span class="post__title">${this._title}</span>
    <time class="post__date">${this._date}</time>
    <div class="post__comments">
      <img
        class="post__comments--icon"
        src="images/comment.svg"
        draggable="false"
        alt="Comments"
      />
      <b class="post__comments--count">${this._comments}</b>
    </div>
    <div class="post__rating">
      <img
        class="post__rating--icon"
        src="images/face-smile.svg"
        draggable="false"
        alt="Comments"
      />
      <b class="post__rating--count">${this._rating}</b>
    </div>
  </header>
  <section class="post__body">
    <div class="post__profile">
      <img
        class="post__avatar"
        src="${this._userAvatar}"
        alt="${this._username}'s avatar"
      />
    </div>
    <div class="post__content">
      ${this._content}
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

    const POST_HEADER = this.querySelector(".post__header");
    const POST_RATING = this.querySelector(".post__rating--count");
    const POST_UPVOTE = this.querySelector(".post__interaction--upvote");
    const POST_DOWNVOTE = this.querySelector(".post__interaction--downvote");
    const POST_COMMENT = this.querySelector(".post__interaction--comment");

    // Make post expandable
    POST_HEADER.addEventListener("click", (event) => {
      if (!event.target.classList.contains("post__username")) {
        this.classList.toggle("post--expanded");
      }
    });
    // Add event listeners to post interactions
    addInteractionListener(
      POST_UPVOTE,
      `/api/posts/${this._ID}/upvote`,
      "POST",
      (data) => {
        if (POST_DOWNVOTE.classList.contains("post__interaction--selected")) {
          POST_DOWNVOTE.classList.toggle("post__interaction--selected");
        }
        POST_UPVOTE.classList.toggle("post__interaction--selected");
        POST_RATING.textContent = data.rating;
      }
    );
    addInteractionListener(
      POST_DOWNVOTE,
      `/api/posts/${this._ID}/downvote`,
      "PUT",
      (data) => {
        if (POST_UPVOTE.classList.contains("post__interaction--selected")) {
          POST_UPVOTE.classList.toggle("post__interaction--selected");
        }
        POST_DOWNVOTE.classList.toggle("post__interaction--selected");
        POST_RATING.textContent = data.rating;
      }
    );
    addInteractionListener(
      POST_COMMENT,
      `/api/posts/${this._ID}/comment`,
      "POST",
      (data) => {
        console.log("Comment");
        // this.querySelector(".post__comments--count").textContent = data.comments.length;
        let commentElement = new Post(
          data.ID,
          data.parentID,
          null,
          data.content,
          data.date,
          [],
          data.rating,
          data.userID,
          data.username,
          data.userAvatar
        );
        this.querySelector(".post__comments").appendChild(commentElement);
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

function addInteractionListener(element, endpoint, method, callback) {
  element.addEventListener("click", () => {
    if (method === "POST" && endpoint.includes("comment")) {
      const commentContent = prompt("Enter your comment:");

      if (!commentContent) {
        return;
      }

      fetch(endpoint, {
        method: method,
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ content: commentContent }),
      })
        .then((response) => response.json())
        .then((data) => {
          callback(data);
        })
        .catch((error) => console.error("Interaction failed:", error));
    } else {
      fetch(endpoint, {
        method: method,
      })
        .then((response) => response.json())
        .then((data) => {
          callback(data);
        })
        .catch((error) => console.error("Interaction failed:", error));
    }
  });
}

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

// DOM loaded
document.addEventListener("DOMContentLoaded", () => {
  populatePosts();
});
