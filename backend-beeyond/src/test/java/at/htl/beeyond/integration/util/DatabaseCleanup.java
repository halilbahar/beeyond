package at.htl.beeyond.integration.util;

import org.eclipse.microprofile.config.ConfigProvider;

import java.sql.*;

public class DatabaseCleanup {
    public static void cleanUp() {
        String databaseUrl = ConfigProvider.getConfig().getValue("quarkus.datasource.jdbc.url", String.class);
        try (
                Connection connection = DriverManager.getConnection(databaseUrl, "beeyond", "beeyond");
                Statement statement = connection.createStatement()
        ) {
            ResultSet tables = statement.executeQuery("SELECT table_name FROM information_schema.tables WHERE table_schema='public'");

            while (tables.next()) {
                statement.execute("TRUNCATE " + tables.getString("table_name") + " CASCADE;");
            }
        } catch (SQLException throwables) {
            throwables.printStackTrace();
        }
    }
}
