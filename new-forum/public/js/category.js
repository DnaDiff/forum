class Category extends HTMLElement {
  constructor(ID, title, posts) {
    super();

    this._ID = ID;
    this._title = title;
    this._posts = posts;

    // Append post to category
    if (this._ID === null) {
      throw new Error("Category must have an ID");
    }
  }

  connectedCallback() {
    // Create HTML elements and set values
    this.id = this._ID;
    this.classList.add("category");
    this.innerHTML = `
    <header class="category__header">
      <span class="category__title">${this._title}</span>
    </header>
    <section class="category__posts">
      <footer class="category__footer"></footer>
    </section>
    `;

    this.querySelector(".category__header").addEventListener("click", () => {
      if (this.querySelector(".category__posts").children.length === 1) {
        this.fetchPosts();
      }
      this.classList.toggle("category--expanded");
    });
  }

  fetchPosts() {
    // Get posts from server
    fetch(`/api/categories/${this._ID}/posts`, {
      method: "GET",
    })
      .then((response) => {
        if (response.ok) {
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
            post.categoryID,
            post.title,
            post.content,
            post.date,
            post.comments,
            post.rating,
            post.userID,
            post.username,
            post.userAvatar
          );
          this.querySelector(".category__posts").insertBefore(
            postElement,
            this.querySelector(".category__footer")
          );
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
      // Create category elements
      Object.values(categories).forEach((category) => {
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
