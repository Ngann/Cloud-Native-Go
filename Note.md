## Anatomy of Cloud Native Stack

* Application Platform: Could Native / Go
* Cluster Orchestrator: Application / Kubernetics
* Cluster Scheduler: Container / Kubernetics / Docker
* Cluster Virtualization: Resources / Docker

Microservice will be deployed in Kubernetics
* Rolling upgrade:
* Scale horizontally:


## What is a Cloud Native Apps?
* Applications designed and implemented to take advantage of the Cloud

## Why Cloud Native Apps?
* Dynamic, won't break and scale aribitrarily
* Hyperscale
* Antifragility only parts of the application will break, as oppose to the whole
* Continious delivery and devops
* Opex is operation expense

## Cloud Native Principles
* Built and composed as microservices
* Package and distributed in cointainers
* Dynamically executed in the cloud

## Challenges and Design Principles
### Challenges
* 

###Design Principles
* Performance: responsive, concurrency, efficiency
* Automation: automate dev tasks and ops task
* Resiliency: fault tolerant and self healing
* Elasticity: Dynamically scale up and down and be reactive
* Delivery: Short round trips and automate delivery
* Diagnosability: Cluster-wide logs , traces and metrics

## Decomposition with microservies
1. Components in the software lifecycle with microservices
    * Design: Complexity, Data Integrity Featured, Decoupled
    * Build: Planning, Knowledge, Development, Integration
    * Run: Release, Deployment, Runtime, Scaling (unique to cloud native app)
2. Ops Component
    * Start interface, internet protocol , diagnosis interface
3. Decomposition tradeoffs
    * increase integration complexity
    * increase infrastructure complexity
    * increase troubleshooting complexity
    
## Intro to Cloud Native Stack
Four layer Architecture

#### Application Platform
Code, contains runtime environment, and API for the application


#### Cluster Operation system

####Orchestrator 
Concern with running whole application, talks to the scheduler to run specific containers and tells it how many instances is needed.

#### Scheduler
Concern with containers, and manages the resources and executes the individual container for those resources

#### Virtualization
Concern with resources in the cluster, decouples everything above it

