**I. Run demo<br/><br/>**
    &nbsp;&nbsp;&nbsp;&nbsp;docker-compose up -d<br/>
    &nbsp;&nbsp;&nbsp;&nbsp;Add following link to postman, there are 3 API for search users, tickets, organizations<br/>
    &nbsp;&nbsp;&nbsp;&nbsp;https://www.getpostman.com/collections/42ecf95a4d33c527d8c6

**II. System describe<br/><br/>**
    &nbsp;&nbsp;&nbsp;&nbsp;Build mongodb with authentication and data sample from tickets.json, users.json, organizations.json<br/>
    &nbsp;&nbsp;&nbsp;&nbsp;Build restful API demo for search users, tickets, organizations<br/>

**III. Optimize note<br/><br/>**
    * Deploy on k8s for optimize: <br/>
        &nbsp;&nbsp;&nbsp;&nbsp;Service discovery and load balancing<br/>
        &nbsp;&nbsp;&nbsp;&nbsp;Automatic rollouts and rollbacks<br/>
        &nbsp;&nbsp;&nbsp;&nbsp;Automated packaging<br/>
        &nbsp;&nbsp;&nbsp;&nbsp;Self-healing<br/>
        &nbsp;&nbsp;&nbsp;&nbsp;Scaling by metric (CPU, Ram, bottleneck, message broker queue...)<br/>
    * API service need be authenticated: <br/>
        &nbsp;&nbsp;&nbsp;&nbsp;Basic-authen (for /api route prefix): <br/>
            &nbsp;&nbsp;&nbsp;&nbsp;Recommend set these info on traefik ingress for k8s<br/>
            &nbsp;&nbsp;&nbsp;&nbsp;Recommend set on Nginx with VPS instance<br/>
        &nbsp;&nbsp;&nbsp;&nbsp;Token (for /public-api route prefix)<br/>
            &nbsp;&nbsp;&nbsp;&nbsp;JWT tokens<br/>
            &nbsp;&nbsp;&nbsp;&nbsp;SSO token (oauth2 standard)<br/>
    * Add SSL for dns<br/>
    * Optimize for search speed: <br/>
        &nbsp;&nbsp;&nbsp;&nbsp;Using cache:<br/>
            &nbsp;&nbsp;&nbsp;&nbsp;Using Go-cache on RAM of main-service<br/>
            &nbsp;&nbsp;&nbsp;&nbsp;Using redis as main cache in this demo. We will clear cache by _id when record is updated.<br/>
       &nbsp;&nbsp;&nbsp;&nbsp;Using Elasticsearch: We will store main fields (which common be searched) and _id of each records to Elasticsearch on indices appropriate with collection name. After get these results, we query by {$in:[_ids]} => Very fast because _id is index by default in mongodb.<br/>
    * Set up EFK for get log from all services<br/>
        &nbsp;&nbsp;&nbsp;&nbsp;EFK (Elasticsearch, FluentBit, Kibana) is a perfect logging and searching text log, which very fit with k8s system<br/>
    * Set up Jaeger for tracing data in all services<br/>

