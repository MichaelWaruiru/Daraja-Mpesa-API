Let me explain the challenges devs face using the Mpesa API
    You have to create a token every time you run your code since the token provided expires after some few minutes.
    Your callbackurl have to be live when going live/production using B2C, C2B or any payment method. This is where you add your domain name/url of your website.

In production,the URLs are provided by Mpesa so you don't to worry about that. In some cases the register URL may fail to work in production yet sandbox works so at that particular time you better ask help from Daraja support team.
   
