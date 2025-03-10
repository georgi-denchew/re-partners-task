const myEnv ={
    apiURL: "http://localhost:3000"
} 


console.log("PROD_MODE", import.meta.env.PROD)
if (import.meta.env.PROD) {
    // todo: this shouldn't be hardcoded
    myEnv.apiURL = "http://3.71.13.28:3000"
}

console.log('API_URL', myEnv.apiURL)
export default myEnv
