class TimeLimitedCache {
  cache = new Map();

  timeoutIds = new Map();

  set(key, value, duration) {
    clearTimeout(this.timeoutIds.get(key));

    const exists = this.cache.has(key);

    this.cache.set(key, value);

    const id = setTimeout(() => {
      this.cache.delete(key);
    }, duration);

    this.timeoutIds.set(key, id);

    return exists;
  }

  get(key) {
    return (this.cache.has(key) && this.cache.get(key)) || -1;
  }

  count() {
    return this.cache.size;
  }
}

const timeLimitedCache = new TimeLimitedCache();
console.log(timeLimitedCache.set(1, 42, 1000)); // false
console.log(timeLimitedCache.get(1)); // 42
console.log(timeLimitedCache.count()); // 1
