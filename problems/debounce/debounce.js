function debounce(fn, t) {
  let id = undefined;

  return function (...args) {
    if (id !== undefined) {
      clearTimeout(id);
    }

    id = setTimeout(() => {
      fn(...args);
    }, t);
  };
}

const log = debounce(console.log, 100);
log(5); // cancelled
log(5); // cancelled
log(5); // // Logged at t=100ms
