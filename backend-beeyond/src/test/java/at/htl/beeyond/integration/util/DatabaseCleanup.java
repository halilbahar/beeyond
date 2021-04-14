package at.htl.beeyond.integration.util;

import org.eclipse.microprofile.config.ConfigProvider;

import java.sql.*;

public class DatabaseCleanup {
    private static final String databaseUrl = ConfigProvider.getConfig().getValue("quarkus.datasource.jdbc.url", String.class);

    public static void cleanUp() {
        try (
                Connection connection = DriverManager.getConnection(databaseUrl, "beeyond", "beeyond");
                Statement statement = connection.createStatement();
        ) {
            ResultSet tables = statement.executeQuery("SELECT table_name FROM information_schema.tables WHERE table_schema='public'");

            while (tables.next()) {
                statement.execute("TRUNCATE " + tables.getString("table_name") + " CASCADE;");
            }
        } catch (SQLException throwables) {
            throwables.printStackTrace();
        }
    }

    public static void insertUsers() {
        try (
                Connection connection = DriverManager.getConnection(databaseUrl, "beeyond", "beeyond");
                Statement statement = connection.createStatement();
        ) {
            statement.execute(
                    "insert into _user (name) values('emina');" +
                            "insert into _user (name) values('stuetz');" +
                            "insert into _user (name) values('moritz');" +
                            "insert into _user (name) values('marc');"
            );
        } catch (SQLException throwables) {
            throwables.printStackTrace();
        }
    }
}
