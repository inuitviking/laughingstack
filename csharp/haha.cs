using System;
using System.Text;

namespace haha
{
    class MainClass
    {
        public static void Main(string[] args)
        {
            var laugh = BuildHahaha();
            Console.WriteLine(laugh);
        }

        private static string BuildHahaha()
        {
            StringBuilder result = new StringBuilder();

            for(int i = 0; i < 3; i++)
            {
                result.Append("ha");
            }

            return result.ToString();
        }
    }
}
