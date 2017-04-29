package com.josebiro.lampcontrol;

import android.content.Context;
import android.content.Intent;
import android.content.SharedPreferences;
import android.os.Bundle;
import android.support.v7.app.AppCompatActivity;
import android.view.View;
import android.widget.Button;
import android.widget.EditText;

/**
 * Created by josebiro on 6/4/2016.
 */
public class SettingsActivity extends AppCompatActivity {
    private SharedPreferences settings;
    private SharedPreferences.Editor settingsEditor;
    private static MainActivity instance;
    private String currentServer;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);

        setContentView(R.layout.settings_activity);

        final EditText serverUrlField = (EditText) findViewById(R.id.serverUrl);
        final Button btnSave = (Button) findViewById(R.id.btnSave);
        final Button btnCancel = (Button) findViewById(R.id.btnCancel);

        instance = MainActivity.get();
        settings = instance.getSharedPreferences("LampControlApp", Context.MODE_PRIVATE);

        currentServer = settings.getString("serverUrl", "http://192.168.86.91:8080");
        serverUrlField.setText(currentServer);

        btnSave.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                // Open the shared preferences and create an editor
                settingsEditor = settings.edit();

                // Save preferences
                String serverUrl = serverUrlField.getText().toString();
                settingsEditor.putString("serverUrl", serverUrl);

                settingsEditor.apply();
                settingsEditor.commit();

                // return to the main view
                Intent act1 = new Intent(v.getContext(), MainActivity.class);
                startActivity(act1);
            }
        });



        btnCancel.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                Intent act1 = new Intent(v.getContext(), MainActivity.class);
                startActivity(act1);
            }
        });
    }
}
