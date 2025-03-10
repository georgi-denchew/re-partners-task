const myEnv ={
    apiURL: "http://localhost:3000"
} 
    

console.log("PROD_MODE", import.meta.env.PROD)
if (import.meta.env.PROD) {
    myEnv.apiURL = "http://18.156.166.29:3000"
}

export default myEnv
