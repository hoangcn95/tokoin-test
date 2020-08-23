**I. Run demo<br/><br/>**
&nbsp;&nbsp;&nbsp;&nbsp;docker-compose up -d<br/>
&nbsp;&nbsp;&nbsp;&nbsp;Add following link to postman, there are 3 API for search users, tickets, organizations<br/>
&nbsp;&nbsp;&nbsp;&nbsp;https://www.getpostman.com/collections/42ecf95a4d33c527d8c6

**II. System describe<br/><br/>**
&nbsp;&nbsp;&nbsp;&nbsp;Build mongodb with authentication and data sample from tickets.json, users.json, organizations.json<br/>
&nbsp;&nbsp;&nbsp;&nbsp;Build restful API demo for search users, tickets, organizations<br/>
&nbsp;&nbsp;&nbsp;&nbsp;There are: 
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;web: contain route, middleware, handler
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;* route: define route inputs and service will be handler. Recommend designing:
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;+ {{host}}/pubic-api/v1/....: Service all service for clients call (Web browser, Mobile app, Desktop app). We can use token for authentication, authorization with this group api.
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;+ {{host}}/api/v1/..........: Service all service for server services call server services. We can basic-authen with this group api. In this group we should devide into two kind: one kind for Restful API, and the other kind for gRPC ( Best option using for optimize speed).
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;+ {{host}}/websocket/.......: Service all socket init connection and transfer data.
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;* middleware: it's design for checking the input context is valid (Have authen, true format, set span for tracing, send logs to EFK...), checking the output is success or not ? Will cache data on RAM of go-service if success...
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;* handler: it's contain all controller business logic functions.
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;model: contain all structs of service
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;repo: contain all connections init, fetch, insert, update, delete query of all DB
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;config: contain all environment variable of three type of real environment. 
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;* Dev: environment for developer test
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;* Uat: environment for QC and user test
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;* Master: environment for production
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;common: contain all libraries which are used between all micro-services. It should be stored all the public repo like: github, gitlab,.. And easily for go get and using
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;utils: contain all common functions,  which high regular using.
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;constants: define all constant, environment variables in program.

**III. Optimize note<br/><br/>**
* Deploy on k8s for optimize: <br/>
&nbsp;&nbsp;&nbsp;&nbsp;Service discovery and load balancing<br/>
&nbsp;&nbsp;&nbsp;&nbsp;Automatic rollouts and rollbacks<br/>
&nbsp;&nbsp;&nbsp;&nbsp;Automated packaging<br/>
&nbsp;&nbsp;&nbsp;&nbsp;Self-healing<br/>
&nbsp;&nbsp;&nbsp;&nbsp;Scaling by metric (CPU, Ram, bottleneck, message broker queue...)<br/>
* API service need be authenticated: <br/>
&nbsp;&nbsp;&nbsp;&nbsp;Basic-authen (for /api route prefix): <br/>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Recommend set these info on traefik ingress for k8s<br/>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Recommend set on Nginx with VPS instance<br/>
&nbsp;&nbsp;&nbsp;&nbsp;Token (for /public-api route prefix)<br/>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;JWT tokens<br/>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;SSO token (oauth2 standard)<br/>
* Add SSL for dns<br/>
* Optimize for search speed: <br/>
&nbsp;&nbsp;&nbsp;&nbsp;Using cache:<br/>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Using Go-cache on RAM of main-service<br/>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Using redis as main cache in this demo. We will clear cache by _id when record is updated.<br/>
&nbsp;&nbsp;&nbsp;&nbsp;Using Elasticsearch: We will store main fields (which common be searched) and _id of each records to Elasticsearch on indices appropriate with collection name. After get these results, we query by {$in:[_ids]} => Very fast because _id is index by default in mongodb.<br/>
* Set up EFK for get log from all services<br/>
&nbsp;&nbsp;&nbsp;&nbsp;EFK (Elasticsearch, FluentBit, Kibana) is a perfect logging and searching text log, which very fit with k8s system<br/>
* Set up Jaeger for tracing data in all services<br/>
