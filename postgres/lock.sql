
/*
*   Fetch all processes which aren't in idle state 
*/
SELECT * FROM pg_stat_activity WHERE state != 'idle';


/*
* Kill a process
*/

SELECT pg_terminate_backend(1234);