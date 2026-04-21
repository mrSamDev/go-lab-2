const obj = new Object({
  name: "main.go",
  age: 30,
  toString: function () {
    return `${this.name} (${this.age} years)`;
  },
});

Object.defineProperty(obj, "toString", {
  value: function () {
    return `${this.name} (${this.age} years)`;
  },
});

console.log(obj.toString());
