class Login{

    constructor(){
        this.config=new AppConfig()
        this.login_path="/user/login"
        this.redirect_uri="/webui/allJobs"
    }

    user_login(email,password,login_url){
        console.log("Executing method to let the user login")
        let body={
            "email":email,
            "password":password
        }
        fetch(login_url,{
            method: 'POST',
            headers:{
                "Accept": 'application.json',
                "Content-Type":'application/json'
            },
            body:JSON.stringify(body)
        })
        .then(res=>res.json())
        .then((data)=>{
            console.log(data.token)
            localStorage.setItem("token",data.token)
            let all_jobs_url=this.config.serverAdress+this.redirect_uri
            window.location.href=all_jobs_url
        })
        .catch(function(e) {
                alert("Unable to login user", e);
            }  
        )
    }
}


document.getElementById('loginform').addEventListener('submit',function(e){
    e.preventDefault()
    let email=document.getElementById('email').value
    let password=document.getElementById('password').value
    console.log(email)
    console.log(password)
    const login=new Login()
    let login_url=login.config.serverAdress+login.login_path
    login.user_login(email,password,login_url)
})