const { WebClient, LogLevel } = require("@slack/web-api");

// WebClient insantiates a client that can call API methods
// When using Bolt, you can use either `app.client` or the `client` passed to listeners.
const client = new WebClient("", {
    // LogLevel can be imported and used to make debugging simpler
    logLevel: LogLevel.DEBUG
});
// You probably want to use a database to store any conversations information ;)
// let conversationsStore = {};

// async function populateConversationStore() {
//     try {
//         // Call the conversations.list method using the WebClient
//         const result = await client.conversations.list({limit: 1000});

//         console.log("length")
//         console.log(result.channels.length)

//         console.log(result.channels)

//         saveConversations(result.channels);
//     }
//     catch (error) {
//         console.error(error);
//     }
// }

// // Put conversations into the JavaScript object
// function saveConversations(conversationsArray) {
//     let conversationId = '';

//     conversationsArray.forEach(function(conversation){
//         // Key conversation info on its unique ID
//         conversationId = conversation["id"];

//         // console.log(conversation)

//         if(conversation['name'] == "devops-support"){
//             console.log(conversationId)
//         }

//         // Store the entire conversation object (you may not need all of the info)
//         conversationsStore[conversationId] = conversation;
//     });


// }

// populateConversationStore(); 




let conversationHistory;
let channelId = 'CE96WL6MQ';
let cicdcounter = 0;

function checkMessage(singleMessage){

    if(singleMessage.indexOf("@ops") != -1){
        cicdcounter +=1 
        console.log("cicd found in message : ", singleMessage)
    }
}

async function getMessages()
{
    try {
        // Call the conversations.history method using WebClient
        const result = await client.conversations.history({
            channel: channelId,
            oldest: "1622523681",
            latest: "1623474081",
            count: "1000"
        });



        conversationHistory = result.messages;

    
        for(let singleMessage of result.messages){
            checkMessage(singleMessage.text)
        }

        // Print results
        console.log(conversationHistory.length + " messages found in " + channelId);

        console.log(cicdcounter)
    } catch (error) {
        console.error(error);
    }
}
getMessages()
