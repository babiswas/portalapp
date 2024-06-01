async function user_login(data){
    try{
        const response=await fetch("http://localhost:3000/api/login",{
            method:"POST",
            headers:{
                "Content-Type": "application/json",
            },
            body:JSON.stringify(data)
        })
        const result = await response.json();
        token_obj=JSON.parse(result)
        localStorage.setItem("token",token_obj.token)
    }catch(error){
        console.error("Error:",error)
    }
}

async function create_new_user(data){
    try{
        const response=await fetch("http://localhost:3000/api/addUser",{
            method:"POST",
            headers:{
                "Content-Type": "application/json",
            },
            body:JSON.stringify(data)
        })
        const result = await response.json();
        user_obj=JSON.parse(result)
        console.log(user_obj)
    }catch(error){
        console.error("Error:",error)
    }
}

async function create_jenkins_job(data){
    try{
        const response=await fetch("http://localhost:3000/api/addUser",{
            method:"POST",
            headers:{
                "Content-Type": "application/json",
            },
            body:JSON.stringify(data)
        })
        const result = await response.json();
        jenkins_job_obj=JSON.parse(result)
        console.log(jenkins_job_obj)
    }catch(error){
        console.error("Error:",error)
    }
}



