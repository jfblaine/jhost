package net.jblaine.jhost.rest;

import javax.ws.rs.ApplicationPath;
import javax.ws.rs.core.Application;
import java.util.concurrent.atomic.AtomicBoolean;

/**
 * Resource entry point
 */
@ApplicationPath("/api")
public class ApplicationConfig extends Application {
    static final AtomicBoolean IS_ALIVE = new AtomicBoolean(true);
}


