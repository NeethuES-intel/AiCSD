# Security Considerations

This guide will cover security considerations for the AiCSD Project.

1. This Project is an Open Source Sample intended to be used as a potential use case/functionality, "Art of the Possible”. 
1. This is not intended to be a full solution; hence it deviates from the Intel SDLe Policy for Open Source Release of not meeting Production SDL. However, it is sufficient to be released as an Open Source Sample.
1. This Open Source Sample is not intended to be used in a Production environment. Changes are required to improve security for a production environment.

## Security Implementation

This system is secured using the secure version of [EdgeX Foundry](https://docs.edgexfoundry.org/2.3/security/Ch-Security/).
The EdgeX message bus is a secure Redis implementation. 
In a single system setup, all services in the solution are configured as [Add-on application services](https://docs.edgexfoundry.org/2.3/security/Ch-Configuring-Add-On-Services/) to Secure EdgeX.
In a two system setup, the Gateway services are configured as add-on services, while the OEM services need to connect using SSH Tunneling.
This implementation is an extension of the EdgeX example for [Remote Device Services in Secure Mode](https://docs.edgexfoundry.org/2.3/security/Ch-RemoteDeviceServices/).

## Additional steps to secure this Open Source Sample

The following list covers the additional features that need to be implemented to secure this open Source Sample:

### OEM
1. Need to add an API gateway such as Kong to control access to OEM endpoints.
1. Communications inside the OEM are not internally secure. Nothing prevents services from being called on localhost.

### Gateway
1. User Authentication needs to be implemented for the Web UI running on the Gateway.
1. Communications inside the gateway are not internally secure. Kong provides an API Gateway for external callers, but does not prevent services from being called on localhost.
2. The external MQTT broker for exporting results is not secure and should be replaced with a secure message broker.
1. The Web UI operates over http not https meaning that network communications are not encrypted.

### Other
1. For data privacy guidelines, do not insert any personal identification content in job attributes like filename etc. 
1. Key pairs generated by ssh-keygen (4096-bit RSA keys) should be password protected.
2. Key pairs should be unique to each set of systems and rotated at least every 3 years. 
1. There is currently no access control system in this project.
2. Files should be encrypted and transferred over https instead of http.
2. Implement TLS for Kong, Vault, and Redis (pub/sub) or a zero trust networking solution like OpenZiti in order to prevent the use of the secure port forwarding for communication between the OEM and Gateway.

## AiCSD Crash Mitigation

1. The [Monitoring](../monitoring/overview.md), [Log Analytics](../log-analytics/overview.md) and [Portainer](../getting-started/troubleshooting.md) features help monitor any failed/crashed docker containers.
1. If either microservices fail, then end to end ML pipeline execution will fail.
1. Hence the specific failed docker service can be restarted via [Portainer](../getting-started/troubleshooting.md). If the issue still persists, then all the docker containers can be restarted.
1. If the issue is still not resolved, run the following commands to remove all docker containers, images, volume, and image files. Archived input and output files are available in the **archive** folder on the Gateway. If required, files can be copied and then deleted from the OEM and Gateway file systems.

    ```bash
     make down clean-volumes clean-images clean-files 
    ```

1. Users can refer to [Troubleshooting Services](../getting-started/troubleshooting.md) and [Troubleshooting Tools](../troubleshooting-tools.md) sections to debug and resolve any issues encountered.