# Minimal Example


1. Launch your Fly App

    ```
    flyctl launch
    ```

1. Upload the secrets

    ```
    flyhelper secrets push
    ```

1. Deploy your app

    ```
    flyctl deploy
    ```

1. Once built and deployed, you should be able to access your app and see the `secret.hello.txt` file listed.

    ```
    echo "https://$(flyctl info --host)"
    ```

1. You can try edit the `secret.txt` file localy and push the secrets again. The app will automatically redeploy.

   ```
   echo "This is the real secret!" > secret.txt
   flyhelper secrets push
   ```

    You should see updated content served once the deployment is finished.

