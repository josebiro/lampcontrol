package com.josebiro.lampcontrol;

import android.net.Uri;
import android.os.AsyncTask;
import android.util.Log;
import android.view.View;
import android.widget.Button;
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

    String lampServer = "http://192.168.86.91:8080";
    String action;
    Button btnLampOn;
    Button btnLampOff;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        // ATTENTION: This was auto-generated to implement the App Indexing API.
        // See https://g.co/AppIndexing/AndroidStudio for more information.
        client = new GoogleApiClient.Builder(this).addApi(AppIndex.API).build();

        btnLampOn = (Button) findViewById(R.id.btn_lamp_on);
        btnLampOff = (Button) findViewById(R.id.btn_lamp_off);

        btnLampOn.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                new LampWhiteOn().execute();
            }
        });
        btnLampOff.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                new LampOff().execute();
            }
        });

    }

    private class LampWhiteOn extends AsyncTask<Void, Void, String> {

        @Override
        protected String doInBackground(Void... params) {
            Log.i("LampWhiteOn", "Entering LampWhiteOn method.");
            HttpURLConnection c = null;
            BufferedReader r = null;
            String response;
            JSONObject action;
            // Send the lamp on command (white light)
            try {
                action = new JSONObject();
                action.put("colorname", "white");
                Log.i("LampWhiteOn", action.toString());
            } catch (JSONException e) {
                Log.e("LampOn", e.getLocalizedMessage());
                return null;
            }
            try {
                byte[] postData = action.toString().getBytes("UTF-8");
                int postDataLength = postData.length;
                Log.i("LampOn", "Content Length is:" + postDataLength);

                URL url = new URL(lampServer + "/setcolor");
                Log.i("LampWhiteOn", url.toString());

                c = (HttpURLConnection) url.openConnection();
                c.setInstanceFollowRedirects(false);
                c.setRequestMethod("POST");
                c.setRequestProperty("X-Custom-Header", "lampcontrol");
                c.setRequestProperty("Content-Type", "application/json; charset=UTF-8");
                c.setRequestProperty("Content-Length", Integer.toString(postDataLength));
                c.setUseCaches(false);
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
                    Log.i("LampOn", "No data returned");
                }
                int status = c.getResponseCode();
                response = buffer.toString();
                Log.i("LampOn", "Response:" + status + response);
            }
            catch (IOException e) {
                Log.i("[ERROR]LampWhiteOn", e.getLocalizedMessage());
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
                        Log.e("LampOn", "Error Closing Stream", e);
                    }
                }
            }
            return response;
        }
    }

    private class LampOff extends AsyncTask<Void, Void, String> {
        @Override
        protected String doInBackground(Void... params) {
            Log.i("LampOff", "Entering LampOff method.");
            HttpURLConnection c = null;
            BufferedReader r = null;
            String response;
            JSONObject action;
            // Send the lamp on command (white light)
            try {
                action = new JSONObject();
                action.put("colorname", "off");
                Log.i("LampOff", action.toString());
            } catch (JSONException e) {
                Log.e("LampOff", e.getLocalizedMessage());
                return null;
            }
            try {
                byte[] postData = action.toString().getBytes("UTF-8");
                int postDataLength = postData.length;

                URL url = new URL(lampServer + "/setcolor");
                Log.i("LampOff", url.toString());

                c = (HttpURLConnection) url.openConnection();
                c.setInstanceFollowRedirects(false);
                c.setRequestMethod("POST");
                c.setRequestProperty("X-Custom-Header", "lampcontrol");
                c.setRequestProperty("Content-Type", "application/json");
                c.setRequestProperty("charset", "utf=8");
                c.setRequestProperty("Content-Length", Integer.toString(postDataLength));
                c.setUseCaches(false);
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
                    Log.i("LampOff", "No data returned");
                }
                int status = c.getResponseCode();
                response = buffer.toString();
                Log.i("LampOff", "Response:" + status + response);
            }
            catch (IOException e) {
                Log.i("[ERROR]LampOff", e.getLocalizedMessage());
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
                        Log.e("LampOff", "Error Closing Stream", e);
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
