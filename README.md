I. Run demo<br/><br/>
    docker-compose up -d<br/>
    Add following link to postman, there are 3 API for search users, tickets, organizations<br/>
    https://www.getpostman.com/collections/42ecf95a4d33c527d8c6

II. System describe<br/><br/>
    _ Build mongodb with authentication and data sample from tickets.json, users.json, organizations.json<br/>
    _ Build restful API demo for search users, tickets, organizations<br/>

III. Optimize note<br/><br/>
    * Deploy on k8s for optimize: <br/>
        Service discovery and load balancing<br/>
        Automatic rollouts and rollbacks<br/>
        Automated packaging<br/>
        Self-healing<br/>
        Scaling by metric (CPU, Ram, bottleneck, message broker queue...)<br/>
    * API service need be authenticated: <br/>
        Basic-authen (for /api route prefix): <br/>
            Recommend set these info on traefik ingress for k8s<br/>
            Recommend set on Nginx with VPS instance<br/>
        Token (for /public-api route prefix)<br/>
            JWT tokens<br/>
            SSO token (oauth2 standard)<br/>
    * Add SSL for dns<br/>
    * Optimize for search speed: <br/>
        Using cache:<br/>
            Using Go-cache on RAM of main-service<br/>
            Using redis as main cache in this demo. We will clear cache by _id when record is updated.<br/>
        Using Elasticsearch: We will store main fields (which common be searched) and _id of each records to Elasticsearch on indices appropriate with collection name. After get these results, we query by {$in:[_ids]} => Very fast because _id is index by default in mongodb.<br/>
    * Set up EFK for get log from all services<br/>
        EFK (Elasticsearch, FluentBit, Kibana) is a perfect logging and searching text log, which very fit with k8s system<br/>
    * Set up Jaeger for tracing data in all services<br/>

