package com.josebiro.lampcontrol;

import android.content.Context;
import android.content.Intent;
import android.content.SharedPreferences;
import android.net.Uri;
import android.os.AsyncTask;
import android.util.Log;
import android.view.View;
import android.widget.Button;
import android.widget.CheckBox;
import android.widget.ImageButton;
import android.widget.SeekBar;
import android.widget.Toast;

import java.io.BufferedInputStream;
import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.io.OutputStream;
import java.net.MalformedURLException;
import java.net.ProtocolException;
import java.net.URL;
import java.net.HttpURLConnection;
import java.nio.charset.StandardCharsets;

import android.graphics.Color;
import com.rtugeek.android.colorseekbar.ColorSeekBar;

import org.json.JSONException;
import org.json.JSONObject;

import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;

import com.google.android.gms.appindexing.Action;
import com.google.android.gms.appindexing.AppIndex;
import com.google.android.gms.common.api.GoogleApiClient;

public class MainActivity extends AppCompatActivity {

    /**
     * ATTENTION: This was auto-generated to implement the App Indexing API.
     * See https://g.co/AppIndexing/AndroidStudio for more information.
     */
    private GoogleApiClient client;
    private static MainActivity instance;

    private SharedPreferences settings;
    private String serverUrl;

    private String action;
    private Button btnLampOn;
    private Button btnLampOff;
    private ImageButton btnSettings;

    private Button btnRed;
    private Button btnBlue;
    private Button btnGreen;
    private Button btnOrange;
    private Button btnIndigo;
    private Button btnWhite;
    private Button btnYellow;
    private Button btnViolet;
    private Button btnPurple;

    public String color;

    public static MainActivity get() {
        return instance;
    }

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        MainActivity.instance = this;

        this.settings = getSharedPreferences("LampControlApp", Context.MODE_PRIVATE);

        serverUrl = settings.getString("serverUrl", "http://192.168.86.91:8080");

        setContentView(R.layout.activity_main);
        // ATTENTION: This was auto-generated to implement the App Indexing API.
        // See https://g.co/AppIndexing/AndroidStudio for more information.
        client = new GoogleApiClient.Builder(this).addApi(AppIndex.API).build();

        btnLampOn = (Button) findViewById(R.id.btn_lamp_on);
        btnLampOff = (Button) findViewById(R.id.btn_lamp_off);
        btnSettings = (ImageButton) findViewById(R.id.btn_settings);

        btnRed = (Button) findViewById(R.id.btn_red);
        btnBlue = (Button) findViewById(R.id.btn_blue);
        btnGreen = (Button) findViewById(R.id.btn_green);
        btnOrange = (Button) findViewById(R.id.btn_orange);
        btnIndigo = (Button) findViewById(R.id.btn_indigo);
        btnWhite = (Button) findViewById(R.id.btn_white);
        btnYellow = (Button) findViewById(R.id.btn_yellow);
        btnViolet = (Button) findViewById(R.id.btn_violet);
        btnPurple = (Button) findViewById(R.id.btn_purple);

        final ColorSeekBar colorSeekBar = (ColorSeekBar) findViewById(R.id.colorSlider);
        //final CheckBox showAlphaCheckBox = (CheckBox) findViewById(R.id.checkBox);
        //final SeekBar barHeightSeekBar = (SeekBar) findViewById(R.id.seekBar);
        //final SeekBar thumbHeightSeekBar = (SeekBar) findViewById(R.id.seekBar2);

        btnLampOn.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                new SetColor().execute("white");
            }
        });
        btnLampOff.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                new SetColor().execute("off");
            }
        });
        btnSettings.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                Intent settingsActivity = new Intent(v.getContext(),SettingsActivity.class);
                startActivity(settingsActivity);
            }
        });
        btnRed.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                new SetColor().execute("red");
            }
        });
        btnBlue.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                new SetColor().execute("blue");
            }
        });
        btnGreen.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                new SetColor().execute("green");
            }
        });
        btnOrange.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                new SetColor().execute("orange");
            }
        });
        btnIndigo.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                new SetColor().execute("indigo");
            }
        });
        btnWhite.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                new SetColor().execute("white");
            }
        });
        btnYellow.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                new SetColor().execute("yellow");
            }
        });
        btnViolet.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                new SetColor().execute("violet");
            }
        });
        btnPurple.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                new SetColor().execute("purple");
            }
        });
        colorSeekBar.setColors(R.array.material_colors);
        colorSeekBar.setBarHeight(10);
        colorSeekBar.setThumbHeight(50);
        colorSeekBar.setOnColorChangeListener(new ColorSeekBar.OnColorChangeListener() {
            @Override
            public void onColorChangeListener(int colorBarValue, int alphaBarValue, int color) {
                String c = "#" + Integer.toHexString(color).substring(2);
                Log.i("ColorSeekBar","color: " + c);
                new SetColorHex().execute(c);
            }
        });
    }

    private class SetColor extends AsyncTask<String, Void, String> {

        @Override
        protected String doInBackground(String... params) {
            Log.i("SetColor", "Entering SetColor method.");
            Log.i("SetColor", "Setting color to " + params[0]);
            HttpURLConnection c = null;
            BufferedReader r = null;
            String response;
            JSONObject action;
            // Send the lamp on command (white light)
            try {
                action = new JSONObject();
                action.put("colorname", params[0]);
                Log.i("SetColor", action.toString());
            } catch (JSONException e) {
                Log.e("SetColor", e.getLocalizedMessage());
                return null;
            }
            try {
                byte[] postData = action.toString().getBytes("UTF-8");
                int postDataLength = postData.length;
                Log.i("SetColor", "Content Length is:" + postDataLength);

                URL url = new URL(serverUrl + "/setcolor");
                Log.i("SetColor", url.toString());

                c = (HttpURLConnection) url.openConnection();
                c.setInstanceFollowRedirects(false);
                c.setRequestMethod("POST");
                c.setRequestProperty("X-Custom-Header", "lampcontrol");
                c.setRequestProperty("Content-Type", "application/json; charset=UTF-8");
                c.setRequestProperty("Content-Length", Integer.toString(postDataLength));
                c.setUseCaches(false);
                c.setConnectTimeout(500);
                OutputStream os = c.getOutputStream();
                os.write(postData);
                os.close();

                InputStream input = new BufferedInputStream(c.getInputStream());
                StringBuffer buffer = new StringBuffer();
                if (input == null) {
                    // nothing to do
                    return null;
                }
                r = new BufferedReader(new InputStreamReader(input));

                String line;
                while ((line = r.readLine()) != null) {
                    buffer.append(line + "\n");
                }
                if (buffer.length() == 0) {
                    Log.i("SetColor", "No data returned");
                }
                int status = c.getResponseCode();
                response = buffer.toString();
                Log.i("SetColor", "Response:" + status + response);
            }
            catch (IOException e) {
                Log.i("[ERROR]SetColor", e.getLocalizedMessage());
                return null;
            }
            finally {
                if (c != null) {
                    c.disconnect();
                }
                if (r != null) {
                    try {
                        r.close();
                    } catch (final IOException e) {
                        Log.e("SetColor", "Error Closing Stream", e);
                    }
                }
            }
            return response;
        }
    }

    private class SetColorHex extends AsyncTask<String, Void, String> {

        @Override
        protected String doInBackground(String... params) {
            Log.i("SetColor", "Entering SetColor method.");
            Log.i("SetColor", "Setting color to " + params[0]);
            HttpURLConnection c = null;
            BufferedReader r = null;
            String response;
            JSONObject action;
            // Send the lamp on command (white light)
            try {
                action = new JSONObject();
                action.put("colorhex", params[0]);
                Log.i("SetColor", action.toString());
            } catch (JSONException e) {
                Log.e("SetColor", e.getLocalizedMessage());
                return null;
            }
            try {
                byte[] postData = action.toString().getBytes("UTF-8");
                int postDataLength = postData.length;
                Log.i("SetColor", "Content Length is:" + postDataLength);

                URL url = new URL(serverUrl + "/setcolor");
                Log.i("SetColor", url.toString());

                c = (HttpURLConnection) url.openConnection();
                c.setInstanceFollowRedirects(false);
                c.setRequestMethod("POST");
                c.setRequestProperty("X-Custom-Header", "lampcontrol");
                c.setRequestProperty("Content-Type", "application/json; charset=UTF-8");
                c.setRequestProperty("Content-Length", Integer.toString(postDataLength));
                c.setUseCaches(false);
                c.setConnectTimeout(500);
                OutputStream os = c.getOutputStream();
                os.write(postData);
                os.close();

                InputStream input = new BufferedInputStream(c.getInputStream());
                StringBuffer buffer = new StringBuffer();
                if (input == null) {
                    // nothing to do
                    return null;
                }
                r = new BufferedReader(new InputStreamReader(input));

                String line;
                while ((line = r.readLine()) != null) {
                    buffer.append(line + "\n");
                }
                if (buffer.length() == 0) {
                    Log.i("SetColor", "No data returned");
                }
                int status = c.getResponseCode();
                response = buffer.toString();
                Log.i("SetColor", "Response:" + status + response);
            }
            catch (IOException e) {
                Log.i("[ERROR]SetColor", e.getLocalizedMessage());
                return null;
            }
            finally {
                if (c != null) {
                    c.disconnect();
                }
                if (r != null) {
                    try {
                        r.close();
                    } catch (final IOException e) {
                        Log.e("SetColor", "Error Closing Stream", e);
                    }
                }
            }
            return response;
        }
    }


    @Override
    public void onStart() {
        super.onStart();

        // ATTENTION: This was auto-generated to implement the App Indexing API.
        // See https://g.co/AppIndexing/AndroidStudio for more information.
        client.connect();
        Action viewAction = Action.newAction(
                Action.TYPE_VIEW, // TODO: choose an action type.
                "Main Page", // TODO: Define a title for the content shown.
                // TODO: If you have web page content that matches this app activity's content,
                // make sure this auto-generated web page URL is correct.
                // Otherwise, set the URL to null.
                Uri.parse("http://host/path"),
                // TODO: Make sure this auto-generated app URL is correct.
                Uri.parse("android-app://com.josebiro.lampcontrol/http/host/path")
        );
        AppIndex.AppIndexApi.start(client, viewAction);
    }

    @Override
    public void onStop() {
        super.onStop();

        // ATTENTION: This was auto-generated to implement the App Indexing API.
        // See https://g.co/AppIndexing/AndroidStudio for more information.
        Action viewAction = Action.newAction(
                Action.TYPE_VIEW, // TODO: choose an action type.
                "Main Page", // TODO: Define a title for the content shown.
                // TODO: If you have web page content that matches this app activity's content,
                // make sure this auto-generated web page URL is correct.
                // Otherwise, set the URL to null.
                Uri.parse("http://host/path"),
                // TODO: Make sure this auto-generated app URL is correct.
                Uri.parse("android-app://com.josebiro.lampcontrol/http/host/path")
        );
        AppIndex.AppIndexApi.end(client, viewAction);
        client.disconnect();
    }
}
