interface IGitObj {
    username: string,
    url: string
}

const obj = <IGitObj>{
    username: "hixb",
    url: "https://github.com/hixb"
}

console.log(obj)