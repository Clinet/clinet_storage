package storage

import (
	"fmt"
	"io/ioutil"

	"github.com/JoshuaDoes/json"
)

//Storage is designed for stateless interactions, to set it and forget it until it's needed again
type Storage struct {
	Configs  map[string]*StorageObject `json:"configs,omitempty"`
	Channels map[string]*StorageObject `json:"channels,omitempty"`
	Messages map[string]*StorageObject `json:"messages,omitempty"`
	Servers  map[string]*StorageObject `json:"servers,omitempty"`
	Users    map[string]*StorageObject `json:"users,omitempty"`
	path string
}

func (s *Storage) LoadFrom(state string) error {
	if s == nil {
		return fmt.Errorf("unable to load state %s into nil storage", state)
	}

	s.path = "states/" + state + ".json"

	stateJSON, err := ioutil.ReadFile(s.path)
	if err != nil {
		return s.Reset()
	}

	if err := json.Unmarshal(stateJSON, s); err != nil {
		return s.Reset()
	}

	return nil
}

func (s *Storage) Reset() error {
	s = &Storage{
		Configs: make(map[string]*StorageObject),
		Channels: make(map[string]*StorageObject),
		Messages: make(map[string]*StorageObject),
		Servers: make(map[string]*StorageObject),
		Users: make(map[string]*StorageObject),
		path: s.path,
	}
	return s.Save()
}

func (s *Storage) Save() error {
	stateJSON, err := json.Marshal(s, true)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(s.path, stateJSON, 0644)
}

type StorageObject struct {
	Data map[string]interface{} `json:"data,omitempty"`
}
func (so *StorageObject) Get(key string) (interface{}, error) {
	if data, exists := so.Data[key]; exists {
		return data, nil
	}
	return nil, fmt.Errorf("invalid key: %s", key)
}
func (so *StorageObject) Set(key string, val interface{}) {
	so.Data[key] = val
}
func (so *StorageObject) Del(key string) {
	delete(so.Data, key)
}

func (s *Storage) ConfigGet(extraID, key string) (interface{}, error) {
	if s.Configs == nil {
		s.Configs = make(map[string]*StorageObject)
	}
	if _, exists := s.Configs[extraID]; exists {
		val, err := s.Configs[extraID].Get(key)
		if err != nil {
			return nil, err
		}
		return val, nil
	}
	return nil, fmt.Errorf("invalid config: %s:%s", extraID, key)
}
func (s *Storage) ConfigSet(extraID, key string, val interface{}) {
	if s.Configs == nil {
		s.Configs = make(map[string]*StorageObject)
	}
	if _, exists := s.Configs[extraID]; !exists {
		s.Configs[extraID] = &StorageObject{
			Data: make(map[string]interface{}),
		}
	}
	s.Configs[extraID].Set(key, val)
	s.Save()
}
func (s *Storage) ConfigDel(extraID, key string) {
	if s.Configs == nil {
		s.Configs = make(map[string]*StorageObject)
	}
	if _, exists := s.Configs[extraID]; exists {
		s.Configs[extraID].Del(key)
	}
}

func (s *Storage) ChannelGet(channelID, key string) (interface{}, error) {
	if s.Channels == nil {
		s.Channels = make(map[string]*StorageObject)
	}
	if _, exists := s.Channels[channelID]; exists {
		return s.Channels[channelID].Get(key)
	}
	return nil, fmt.Errorf("invalid channel: %s:%s", channelID, key)
}
func (s *Storage) ChannelSet(channelID, key string, val interface{}) {
	if s.Channels == nil {
		s.Channels = make(map[string]*StorageObject)
	}
	if _, exists := s.Channels[channelID]; !exists {
		s.Channels[channelID] = &StorageObject{
			Data: make(map[string]interface{}),
		}
	}
	s.Channels[channelID].Set(key, val)
	s.Save()
}
func (s *Storage) ChannelDel(channelID, key string) {
	if s.Channels == nil {
		s.Channels = make(map[string]*StorageObject)
	}
	if _, exists := s.Channels[channelID]; exists {
		s.Channels[channelID].Del(key)
	}
}

func (s *Storage) MessageGet(messageID, key string) (interface{}, error) {
	if s.Messages == nil {
		s.Messages = make(map[string]*StorageObject)
	}
	if _, exists := s.Messages[messageID]; exists {
		return s.Messages[messageID].Get(key)
	}
	return nil, fmt.Errorf("invalid message: %s:%s", messageID, key)
}
func (s *Storage) MessageSet(messageID, key string, val interface{}) {
	if s.Messages == nil {
		s.Messages = make(map[string]*StorageObject)
	}
	if _, exists := s.Messages[messageID]; !exists {
		s.Messages[messageID] = &StorageObject{
			Data: make(map[string]interface{}),
		}
	}
	s.Messages[messageID].Set(key, val)
	s.Save()
}
func (s *Storage) MessageDel(messageID, key string) {
	if s.Messages == nil {
		s.Messages = make(map[string]*StorageObject)
	}
	if _, exists := s.Messages[messageID]; exists {
		s.Messages[messageID].Del(key)
	}
}

func (s *Storage) ServerGet(serverID, key string) (interface{}, error) {
	if s.Servers == nil {
		s.Servers = make(map[string]*StorageObject)
	}
	if _, exists := s.Servers[serverID]; exists {
		return s.Servers[serverID].Get(key)
	}
	return nil, fmt.Errorf("invalid server: %s:%s", serverID, key)
}
func (s *Storage) ServerSet(serverID, key string, val interface{}) {
	if s.Servers == nil {
		s.Servers = make(map[string]*StorageObject)
	}
	if _, exists := s.Servers[serverID]; !exists {
		s.Servers[serverID] = &StorageObject{
			Data: make(map[string]interface{}),
		}
	}
	s.Servers[serverID].Set(key, val)
	s.Save()
}
func (s *Storage) ServerDel(serverID, key string) {
	if s.Servers == nil {
		s.Servers = make(map[string]*StorageObject)
	}
	if _, exists := s.Servers[serverID]; exists {
		s.Servers[serverID].Del(key)
	}
}

func (s *Storage) UserGet(userID, key string) (interface{}, error) {
	if s.Users == nil {
		s.Users = make(map[string]*StorageObject)
	}
	if _, exists := s.Users[userID]; exists {
		return s.Users[userID].Get(key)
	}
	return nil, fmt.Errorf("invalid user: %s:%s", userID, key)
}
func (s *Storage) UserSet(userID, key string, val interface{}) {
	if s.Users == nil {
		s.Users = make(map[string]*StorageObject)
	}
	if _, exists := s.Users[userID]; !exists {
		s.Users[userID] = &StorageObject{
			Data: make(map[string]interface{}),
		}
	}
	s.Users[userID].Set(key, val)
	s.Save()
}
func (s *Storage) UserDel(userID, key string) {
	if s.Users == nil {
		s.Users = make(map[string]*StorageObject)
	}
	if _, exists := s.Users[userID]; exists {
		s.Users[userID].Del(key)
	}
}