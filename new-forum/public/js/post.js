class Post extends HTMLElement {
  constructor(
    ID,
    categoryID,
    title,
    content,
    date,
    rating,
    userID,
    username,
    userAvatar
  ) {
    super();

    this._ID = ID;
    this._categoryID = categoryID;
    this._title = title;
    this._content = content;
    this._date = date;
    this._rating = rating;
    this._userID = userID;
    this._username = username;
    this._userAvatar = userAvatar;

    if (this._categoryID === null) {
      throw new Error("Post must have a categoryID");
    }
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
    `/api/categories/${this._categoryID}/posts/${this._ID}`,
    "POST",
    (data) => {
      if (POST_DOWNVOTE.classList.contains("post__interaction--selected")) {
        POST_DOWNVOTE.classList.toggle("post__interaction--selected");
      }
      POST_UPVOTE.classList.toggle("post__interaction--selected");
      POST_RATING.textContent = data.rating;
    },
    { action: "upvote" }
  );
  addInteractionListener(
    POST_DOWNVOTE,
    `/api/categories/${this._categoryID}/posts/${this._ID}`,
    "PUT",
    (data) => {
      if (POST_UPVOTE.classList.contains("post__interaction--selected")) {
        POST_UPVOTE.classList.toggle("post__interaction--selected");
      }
      POST_DOWNVOTE.classList.toggle("post__interaction--selected");
      POST_RATING.textContent = data.rating;
    },
    { action: "downvote" }
  );
  
  
    // Add event listener to comment icon
    POST_COMMENT.addEventListener("click", () => {
      // Open the comment form modal
      const commentFormModal = document.getElementById("comment-form-modal");
      const commentForm = commentFormModal.querySelector('#comment-form');
      commentForm.setAttribute('data-category-id', this._categoryID);
      commentForm.setAttribute('data-post-id', this._ID);
      commentFormModal.style.display = "block";

     // Close the comment form modal when the close button is clicked
      const closeBtn = commentFormModal.querySelector(".close-btn");
      closeBtn.addEventListener("click", () => {
      commentFormModal.style.display = "none";
      });
    });
  }


  disconnectedCallback() {
    console.log("Post removed");
    // Remove post
  }
}
customElements.define("post-element", Post);

function addInteractionListener(
  element,
  endpoint,
  method,
  callback,
  data = {}
) {
  element.addEventListener("click", () => {
    fetch(endpoint, {
      method: method,
      body: JSON.stringify(data),
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }
        return response.json();
      })
      .then((data) => {
        callback(data);
      })
      .catch((error) => console.error("Interaction failed:", error));
  });
}


// DOM loaded
document.addEventListener("DOMContentLoaded", () => {
  // populatePosts();
  document.getElementById('submit-comment-button').addEventListener('click', (event) => {
    event.preventDefault();
    const comment = document.getElementById('comment-input').value;
    const commentForm = document.getElementById('comment-form');
    const categoryID = commentForm.getAttribute('data-category-id');
    const postID = commentForm.getAttribute('data-post-id');
  
    fetch(`/api/categories/${categoryID}/posts/${postID}/comments`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({content: comment}),
    })
    .then((response) => {
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      console.log(response);
      return response.json();
    })
    .then(data => {
      console.log("i dont come here")
      // Handle the response data
      alert(data.message);
      console.log(data);
    })
    .catch((error) => {
      console.error('Error:', error);
    });
  });  
  document.getElementById('post-form').addEventListener('submit', (event) => {
    event.preventDefault();
    const title = document.getElementById('title').value;
    const content = document.getElementById('content').value;
    const categoryID = document.getElementById('category-select').value;
  
    fetch(`/api/categories/${categoryID}/posts`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ title: title, content: content }),
    })
    .then(response => {
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      return response.json();
    })
    .then(data => {
      // Handle the response data
      alert(data.message);
      console.log(data);
    })
    .catch((error) => {
      console.error('Error:', error);
    });
  });  
});