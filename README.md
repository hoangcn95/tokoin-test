I. Run demo. 
    docker-compose up -d
    Add following link to postman, there are 3 API for search users, tickets, organizations
    https://www.getpostman.com/collections/42ecf95a4d33c527d8c6

II. System describe
    _ Build mongodb with authentication and data sample from tickets.json, users.json, organizations.json
    _ Build restful API demo for search users, tickets, organizations

III. Optimize note
    * Deploy on k8s for optimize: 
        Service discovery and load balancing
        Automatic rollouts and rollbacks
        Automated packaging
        Self-healing
        Scaling by metric (CPU, Ram, bottleneck, message broker queue...)
    * API service need be authenticated: 
        Basic-authen (for /api route prefix): 
            Recommend set these info on traefik ingress for k8s
            Recommend set on Nginx with VPS instance
        Token (for /public-api route prefix)
            JWT tokens
            SSO token (oauth2 standard)
    * Add SSL for dns
    * Optimize for search speed: 
        Using cache:
            Using Go-cache on RAM of main-service
            Using redis as main cache in this demo. We will clear cache by _id when record is updated.
        Using Elasticsearch: We will store main fields (which common be searched) and _id of each records to Elasticsearch on indices appropriate with collection name. After get these results, we query by {$in:[_ids]} => Very fast because _id is index by default in mongodb.
    * Set up EFK for get log from all services
        EFK (Elasticsearch, FluentBit, Kibana) is a perfect logging and searching text log, which very fit with k8s system
    * Set up Jaeger for tracing data in all services

