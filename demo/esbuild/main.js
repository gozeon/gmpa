class A {
    constructor(a) {
        this.a = a
        this.a()
    }

    async a() {
        console.log(this?.b)
    }
}
const a = new A()
