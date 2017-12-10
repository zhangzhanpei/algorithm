/**
 * 给定一个员工数组，包括员工id，员工重要性，以及他的下级员工id，求某个员工的重要性
 * 重要性 = 该员工的重要性 + 递归下级员工的重要性之和
 */
public int getImportance(List<Employee> employees, int id) {
    Map<Integer, Employee> map = new HashMap<>(); //所有员工存到HashMap中
    for (Employee e : employees) {
        map.put(e.id, e);
    }
    return dfs(map, id);
}

public int dfs(Map<Integer, Employee> map, int eid) {
    Employee e = map.get(eid);
    int total = e.importance; //当前员工重要性
    for (int subordinate : e.subordinates) { //递归求下级员工重要性之和
        total += dfs(map, subordinate);
    }
    return total;
}
