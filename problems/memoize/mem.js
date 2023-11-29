function memoize(fn) {
  const cache = new Map();
  return function (...args) {
    const k = args.join(',');
    if (cache.has(k)) {
      return cache.get(k);
    }
    const v = fn(...args);
    cache.set(k, v);
    return v;
  };
}

let callCount = 0;
const memoizedFn = memoize(function (a, b) {
  callCount += 1;
  return a + b;
});
memoizedFn(2, 3); // 5
memoizedFn(2, 3); // 5
console.log(callCount); // 1
