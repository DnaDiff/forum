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
  generateTemplate() {
    return `
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
  }

  connectedCallback() {
    // Create HTML elements and set values
    this.id = this._ID;
    this.classList.add("post");
    this.innerHTML = this.generateTemplate();

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
  element.addEventListener("click", async () => {
    // Make this callback async
    if (method === "POST" && endpoint.includes("comment")) {
      const loggedIn = await isLoggedIn();
      if (!loggedIn) {
        alert("You must be logged in to comment.");
        return;
      }
      const commentFormModal = document.getElementById("comment-form-modal");
      const dim = document.getElementById("dim");
      commentFormModal.style.display = "block";
      dim.style.display = "block";

      const closeBtn = commentFormModal.querySelector(".close-btn");
      closeBtn.addEventListener("click", () => {
        commentFormModal.style.display = "none";
        dim.style.display = "none";
      });

      dim.addEventListener("click", () => {
        commentFormModal.style.display = "none";
        dim.style.display = "none";
      });

      const commentForm = document.getElementById("comment-form");
      commentForm.addEventListener("submit", async (event) => {
        event.preventDefault();

        const commentContent = document.getElementById("comment-content").value;

        if (!commentContent) {
          return;
        }

        try {
          const response = await fetch(endpoint, {
            method: method,
            headers: {
              "Content-type": "application/json",
            },
            body: JSON.stringify({ content: commentContent }),
          });
          const data = await response.json();
          callback(data);

          document.getElementById("comment-content").value = "";
          commentFormModal.style.display = "none";
          dim.style.display = "none";
        } catch (error) {
          console.error("Comment submission failed:", error);
        }
      });

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

  // Add event listeners for opening and closing modals
  const modalButtons = document.querySelectorAll(".modal-btn");
  modalButtons.forEach((button) => {
    button.addEventListener("click", (event) => {
      const modalId = event.target.dataset.modal;
      document.getElementById(modalId).classList.add("modal--visible");
      document.getElementById("dim").classList.add("dim--visible");
    });
  });

  // Add event listeners for closing modals
  const closeButtons = document.querySelectorAll(".close-btn");
  closeButtons.forEach((button) => {
    button.addEventListener("click", (event) => {
      const modal = event.target.closest(".modal");
      modal.classList.remove("modal--visible");
      document.getElementById("dim").classList.remove("dim--visible");
    });
  });

  // Add event listener for comment form submission
  const commentForm = document.getElementById("comment-form");
  commentForm.addEventListener("submit", async (event) => {
    event.preventDefault();

    const postId = commentForm.dataset.postId;
    const commentContent = document.getElementById("comment-content").value;

    if (!commentContent) {
      return;
    }

    fetch(`/api/posts/${postId}/comment`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ content: commentContent }),
    })
      .then((response) => response.json())
      .then((data) => {
        let postElement = document.querySelector(
          `post-element[id="${postId}"]`
        );
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
        postElement
          .querySelector(".post__comments")
          .appendChild(commentElement);
      })
      .catch((error) => console.error("Interaction failed:", error));

    // Close the comment modal and clear the input
    document
      .getElementById("comment-form-modal")
      .classList.remove("modal--visible");
    document.getElementById("dim").classList.remove("dim--visible");
    document.getElementById("comment-content").value = "";
  });
});

async function isLoggedIn() {
  const sessionToken = getCookie("sessionToken");

  if (!sessionToken) {
    return false;
  }

  try {
    const response = await fetch("/api/auth/check", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ token: sessionToken }),
    });

    const data = await response.json();
    if (data.valid) {
      return true;
    } else {
      return false;
    }
  } catch (error) {
    console.error("Error checking session validity:", error);
    return false;
  }
}
