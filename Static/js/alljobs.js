class JenkinsJob{

    constructor(){
        this.config=new AppConfig()
        this.login_path="/user/login"
        this.redirect_uri="/jenkins/alljobs"
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
            localStorage.setItem("token",data.token)
        })
        .catch(alert("Unable to login user"))
    }
}


document.getElementById('loginform').addEventListener('submit',function(e){
    e.preventDefault()
    let email=document.getElementById('email').value
    let password=document.getElementById('password').value
    const login=new Login()
    let login_url=login.config.serverAdress+login.login_path
    let all_jobs_url=login.config.serverAdress+login.redirect_uri
    token=login.user_login(email,password,login_url)
    windows.location.href=all_jobs_url
})