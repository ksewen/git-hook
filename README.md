# Introduction

[DE](README_DE.md)

## The reason for development

Because of my former work, I used an IM tool **Slack**. It supports integration with some events from Gitlab and sends messages in the channels. The function is good when you work together within a team.
You should usually have a few **groups** so that you can communicate with each other about each project if you use WeChatWork for work.
The term **group** is similar to the term **channel** in Slack. A **channel** can be used as a project team in which all members have to focus on one or more projects.
That's why I remembered robots on WeChatWork. A robot reports information from the selected repositories to a group in which all the people who are working on these projects are there. Of course, everyone can join several **channels** to receive the latest news from different projects.

## Features
1. Support at some Gitlab events
* Push Event
* Issue Event
* Merge Request Event
2. Configuration for the selected robot through path variable

# Instruction
1. Ask administrator to create a new robot and get the identity of it.
2. Use “Webhook” under Settings -> Integrations of the Git repository.
```
URL: http://host:port/webhooks/gitlab/${robotId}
Secret Token: chat me
Trigger: Check the appropriate “events” according to your needs
```
3. Enjoy it!
