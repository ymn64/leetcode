// NOTE: Any valid solution (even using Promise.all) would fail test case 27
// with a 'Time Limit Exceeded' error. This is a bug.
// TODO: Submit again when they fix this bug.

// then/catch
function promiseAll(functions) {
  return new Promise((resolve, reject) => {
    const results = [];
    const total = functions.length;
    let completed = 0;

    functions.forEach((func, i) => {
      func()
        .then((result) => {
          results[i] = result;
          completed++;

          if (completed === total) {
            resolve(results);
          }
        })
        .catch((error) => {
          reject(error);
        });
    });
  });
}

// async/await
async function promiseAll2(functions) {
  return new Promise((resolve, reject) => {
    const results = [];
    const total = functions.length;
    let completed = 0;

    try {
      functions.forEach(async (func, i) => {
        results[i] = await func();
        completed++;

        if (completed === total) {
          resolve(results);
        }
      });
    } catch (error) {
      reject(error);
    }
  });
}

const promise = promiseAll([() => new Promise((res) => res(42))]);
promise.then(console.log); // [42]

// Test case 27
const promise2 = promiseAll2([
  () => new Promise((resolve) => setTimeout(() => resolve(1), 100)),
  () => new Promise((resolve) => setTimeout(() => resolve(2), 80)),
  () => new Promise((resolve) => setTimeout(() => resolve(3), 50)),
  () => new Promise((resolve) => setTimeout(() => resolve(4), 20)),
  () => new Promise((resolve) => setTimeout(() => resolve(5), 280)),
  () => new Promise((resolve) => setTimeout(() => resolve(6), 250)),
  () => new Promise((resolve) => setTimeout(() => resolve(7), 220)),
  () => new Promise((resolve) => setTimeout(() => resolve(8), 130)),
  () => new Promise((resolve) => setTimeout(() => resolve(9), 160)),
  () => new Promise((resolve) => setTimeout(() => resolve(10), 190)),
]);
promise2.then(console.log);
