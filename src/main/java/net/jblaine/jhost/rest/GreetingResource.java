package net.jblaine.jhost.rest;

import javax.ws.rs.*;
import javax.ws.rs.core.Response;

@Path("/")
public class GreetingResource {

    private static final String template = "Hello, %s!";

    @GET
    @Path("/greeting")
    @Produces("application/json")
    public Greeting greeting(@QueryParam("name") @DefaultValue("World") String name) {
        if (!ApplicationConfig.IS_ALIVE.get()) {
            throw new WebApplicationException(Response.Status.SERVICE_UNAVAILABLE);
        }
        return new Greeting(String.format(template, name));
    }

    @GET
    @Path("/stop")
    public Response stop() {

        if (!ApplicationConfig.IS_ALIVE.get()) {
            throw new WebApplicationException(Response.Status.SERVICE_UNAVAILABLE);
        }

        ApplicationConfig.IS_ALIVE.set(false);
        return Response.ok("killed").build();

    }

}

