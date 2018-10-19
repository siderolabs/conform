const { events, Job, Group } = require("brigadier");

events.on("exec", function (e, project) {
    job = conform(e, project)
    job.run().then(result => {
        console.log(result.toString())
    })
})

events.on("push", function (e, project) {
    job = conform(e, project)
    job.run().then(result => {
        console.log(result.toString())
    })
})

events.on("pull_request", function (e, project) {
    start = notify("pending", `Build ${e.buildID} started`, e, project)
    job = conform(e, project)
    Group.runAll([start, job])
        .then(() => {
            return notify("success", `Build ${e.buildID} passed`, e, project).run()
        }).catch(err => {
            return notify("failure", `Build ${e.buildID} failed`, e, project).run()
        });
})

function conform(e, project) {
    var job = new Job("conform", "golang:1.11.1")

    job.env = {
        // "DOCKER_HOST": "tcp://docker:2375"
        "DOCKER_USERNAME": project.secrets.DOCKER_USERNAME,
        "DOCKER_PASSWORD": project.secrets.DOCKER_PASSWORD,
        "GO111MODULE": "on",
    }

    job.tasks = [
        "apt-get update",
        "apt-get -y install apt-transport-https ca-certificates curl software-properties-common",
        "curl -fsSL https://download.docker.com/linux/debian/gpg | apt-key add -",
        "add-apt-repository \"deb [arch=amd64] https://download.docker.com/linux/debian $(lsb_release -cs) stable\"",
        "apt-get update",
        "apt-get -y install docker-ce=18.06.1~ce~3-0~debian",
        "cd /src",
        "go install .",
        "conform enforce",
        "conform build",
    ]

    job.docker.enabled = true

    // Unit is milliseconds, 900000ms = 15m.
    job.timeout = 900000

    job.host.nodeSelector.set("node-role.kubernetes.io/ci", "")

    return job
}

function notify(state, msg, e, project) {
    const gh = new Job(`notify-${state}`, "technosophos/github-notify:latest")
    gh.env = {
        GH_REPO: project.repo.name,
        GH_STATE: state,
        GH_DESCRIPTION: msg,
        GH_CONTEXT: "brigade",
        GH_TOKEN: project.secrets.GH_TOKEN,
        GH_COMMIT: e.revision.commit,
        GH_TARGET_URL: `https://ci.dev.autonomy.io/builds/${e.buildID}`,
    }
    return gh
}
