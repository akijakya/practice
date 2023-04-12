SELECT employee.name, job.type
FROM employee
LEFT JOIN job.type
ON employee.job_id = job.id;