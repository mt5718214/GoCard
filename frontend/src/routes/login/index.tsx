import { component$ } from '@builder.io/qwik';
import { Link } from '@builder.io/qwik-city';
// import { DocumentHead } from '@builder.io/qwik-city';

export default component$(() => {
  return (
    <>
      <div class="container py-5">
        <form class="w-100">
          <div class="text-center mb-4">
            <h1 class="h3 mb-3 font-weight-normal">Sign In</h1>
          </div>

          <div class="form-label-group mb-2">
            <label for="email">email</label>
            <input
              id="email"
              name="email"
              type="email"
              class="form-control"
              placeholder="email"
              required
            />
          </div>

          <div class="form-label-group mb-3">
            <label for="password">Password</label>
            <input
              id="password"
              name="password"
              type="password"
              class="form-control"
              placeholder="Password"
              required
            />
          </div>

          <button
            class="btn btn-lg btn-primary btn-block mb-3"
            type="submit"
          >
            Submit
          </button>

        <div class="text-center mb-3">
          <p>
            <Link href="/signup">Sign Up</Link>
          </p>
        </div>

        <p class="mt-5 mb-3 text-muted text-center">&copy; 2021-2022</p>
      </form>
    </div >
    </>
  );
});

// export const head: DocumentHead = {
//   title: '',
// };
