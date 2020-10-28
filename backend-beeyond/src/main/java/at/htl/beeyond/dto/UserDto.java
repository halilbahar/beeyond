package at.htl.beeyond.dto;

import at.htl.beeyond.entity.User;

public class UserDto {
    
    private Long id;
    
    private String name;

    public UserDto(Long id, String name) {
        this.id = id;
        this.name = name;
    }

    public UserDto() {
    }

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }
    
    public User map() {
        return User.findById(this.id);
    }

    public static UserDto map(User user) {
        return new UserDto(user.getId(), user.getName());
    }
}
