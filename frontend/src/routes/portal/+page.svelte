<script>
  import Box from '$lib/lunaria/Box.svelte'
</script>

<div class="container">
  <Box>
    <img class="user-avatar" src="https://github.com/retrixe.png" alt="logo" />
    <br />
    <h1>Name</h1>
    <h3>PRN</h3>
    <h4>Course e.g. B.Tech CSE</h4>
    <h4>Semester 4</h4>
  </Box>
</div>

<style>
  .container {
    flex: 1;
    display: flex;
    align-items: center;
    justify-content: center;
    > :global(div) {
      padding: 24px;
    }
  }

  .user-avatar {
    width: 100px;
    aspect-ratio: 1;
    border-radius: 100%;
  }
</style>
