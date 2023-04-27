class Category extends HTMLElement {
  constructor(ID, title, posts) {
    super();

    this._ID = ID;
    this._title = title;
    this._posts = posts;
    this._postIndex = 0; // Last post index fetched

    // Append post to category
    if (this._ID === null) {
      throw new Error("Category must have an ID");
    }
    document.querySelector(`.post-container`).appendChild(this);
  }

  connectedCallback() {
    // Create HTML elements and set values
    this.id = this._ID;
    this.classList.add("category");
    this.innerHTML = `
    <header class="category__header">
      <span class="category__title">${this._title}</span>
    </header>
    <section class="category__posts"></section>
    <footer class="category__footer"></footer>
    `;

    this.addEventListener("click", () => {
      this.classList.toggle("category--expanded");
      this.togglePosts();
    });
  }

  togglePosts() {
    const POSTS = this.querySelector(".category__posts");
    if (POSTS.innerHTML === "") {
      this.fetchPosts();
    } else {
      POSTS.innerHTML = "";
    }
  }

  fetchPosts() {
    // Get posts from server
    console.log(this._ID);
    const QUEUED_POST_IDS = this._posts.slice(
      this._postIndex,
      this._postIndex + 5
    );
    fetch(`/api/categories/${this._ID}/posts`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        postIDs: QUEUED_POST_IDS,
      }),
    })
      .then((response) => {
        if (response.ok) {
          postIndex += QUEUED_POST_IDS.length;
          return response.json();
        } else {
          throw new Error("Could not get posts");
        }
      })
      .then((posts) => {
        // Create post elements
        posts.forEach((post) => {
          const postElement = new Post(
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
          this.querySelector(".category__posts").appendChild(postElement);
        });
      })
      .catch((error) => {
        console.error(error);
      });
  }
}

customElements.define("category-element", Category);

function fetchCategories() {
  // Fetch categories from server
  fetch("/api/categories", {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  })
    .then((response) => {
      if (response.ok) {
        return response.json();
      } else {
        throw new Error("Could not get categories");
      }
    })
    .then((categories) => {
      console.log(categories);
      // Create category elements
      Object.values(categories).forEach((category) => {
        console.log(category);
        const categoryElement = new Category(
          category.ID,
          category.title,
          category.posts
        );
        document.querySelector(".post-container").appendChild(categoryElement);
      });
    })
    .catch((error) => {
      console.error(error);
    });
}

// DOM loaded
document.addEventListener("DOMContentLoaded", () => {
  fetchCategories();
});
