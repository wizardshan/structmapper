package main

// entity
//go:generate go run github.com/wizardshan/structmapper -toName User -fromName User -toPath ./domain -fromPath ./repository/ent  -toVar dom -fromVar u -funcPosition from
//go:generate go run github.com/wizardshan/structmapper -toName Users -fromName Users -toPath ./domain -fromPath ./repository/ent  -toVar dom -fromVar u -funcPosition from
//go:generate go run github.com/wizardshan/structmapper -toName Orders -fromName Orders -toPath ./domain -fromPath ./repository/ent  -toVar dom -fromVar o -funcPosition from
//go:generate go run github.com/wizardshan/structmapper -toName Shop -fromName Shop -toPath ./domain -fromPath ./repository/ent  -toVar dom -fromVar s -funcPosition from
//go:generate go run github.com/wizardshan/structmapper -toName Items -fromName Items -toPath ./domain -fromPath ./repository/ent  -toVar dom -fromVar i -funcPosition from

// response
//go:generate go run github.com/wizardshan/structmapper -toName User -fromName User -toPath ./response  -fromPath ./domain -toVar resp -fromVar dom
//go:generate go run github.com/wizardshan/structmapper -toName Users -fromName Users -toPath ./response  -fromPath ./domain -toVar resp -fromVar dom
//go:generate go run github.com/wizardshan/structmapper -toName Orders -fromName Orders -toPath ./response  -fromPath ./domain -toVar resp -fromVar dom

