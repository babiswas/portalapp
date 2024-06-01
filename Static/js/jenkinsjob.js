class JenkinsJob{

    constructor(){
        this.config=new AppConfig()
        this.jenkins_job_path="/jenkins/allJobs"
    }

    display_all_jobs(jenkins_job_url){
        console.log("Executing method to fetch all jenkins job:")
        const table=document.getElementById("mytable")

        fetch(jenkins_job_url,{
            method: 'GET',
            headers:{
                "Accept": 'application.json',
                "Authorization":"Bearer "+localStorage.getItem("token")
            },
        })
        .then(res=>res.json())
        .then((data)=>{
            let jobs_list=data.jobs
            console.log(jobs_list)
            jobs_list.forEach(element=>{
                let table_row=document.createElement("tr")

                let td1=document.createElement("td")
                let td2=document.createElement("td")
                let td3=document.createElement("td")
                let td4=document.createElement("td")
                let td5=document.createElement("td")

                td1.innerText=element.jobname
                td2.innerText=element.project_name
                td3.innerText=element.feature_name
                td4.innerText=element.created_at
                td5.innerText=element.updated_at

                table_row.appendChild(td1)
                table_row.appendChild(td2)
                table_row.appendChild(td3)
                table_row.appendChild(td4)
                table_row.appendChild(td5)

                table.appendChild(table_row)
            })

        })
        .catch(function(e) {
            console.log(e)
        })
    }
}

console.log("Creating jenkins job object.")
let jenkins_obj=new JenkinsJob()

console.log("Inserting rows inside the table.")
let jenkins_url=jenkins_obj.config.serverAdress+jenkins_obj.jenkins_job_path

jenkins_obj.display_all_jobs(jenkins_url)
